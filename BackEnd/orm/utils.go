package orm

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Column struct {
	Name    string
	DBName  string
	Default string
	Desc    []string
	Primary bool
}

func parseDate(v string) (*time.Time, error) {
	if t, err := time.Parse("2006-01-02", v); err != nil {
		return nil, err
	} else {
		return &t, nil
	}
}

type Insert struct {
	db       *gorm.DB
	columns  []*Column
	callback []func(map[string]interface{})
}

func (this *Insert) DB() *gorm.DB {
	return this.db
}

func (this *Insert) Model(model interface{}) *Insert {
	this.columns = nil
	rt := reflect.TypeOf(model)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	var cols []*Column
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag
		col := &Column{}
		col.Name = field.Name
		col.Desc = strings.Split(tag.Get("desc"), "|")
		for _, gorm := range strings.Split(tag.Get("gorm"), ";") {
			v := strings.Split(gorm, ":")
			k := v[0]
			switch k {
			case "primary_key":
				col.Primary = true
			case "column":
				col.DBName = v[1]
			case "default":
				col.Default = strings.Trim(v[1], "'")
			}
		}
		cols = append(cols, col)
	}
	this.columns = cols
	this.db = DB().Model(model)
	return this
}

func (this *Insert) Statement() *gorm.Statement {
	return this.db.Statement
}

func (this *Insert) SetCallback(callback ...func(map[string]interface{})) {
	this.callback = callback
}

func (this *Insert) Values(values []map[string]interface{}, update bool) error {
	var conflict *clause.OnConflict
	if update {
		var primary []clause.Column
		var update []string
		for _, f := range this.columns {
			if f.Primary {
				primary = append(
					primary,
					clause.Column{Name: f.DBName},
				)
			} else {
				update = append(
					update,
					f.DBName,
				)
			}

		}
		conflict = &clause.OnConflict{
			Columns:   primary,
			DoUpdates: clause.AssignmentColumns(update),
		}
	} else {
		conflict = &clause.OnConflict{
			DoNothing: true,
		}
	}
	defaults := make(map[string]string)
	for _, c := range this.columns {
		defaults[c.Name] = c.Default
	}
	for _, value := range values {
		for k, d := range defaults {
			if v, ok := value[k]; !ok || v == "" {
				value[k] = d
			}
		}
		if this.callback != nil {
			for _, f := range this.callback {
				f(value)
			}
		}
	}
	this.db = this.db.Clauses(conflict).Create(values)
	return this.db.Error
}

type InsertSheet struct {
	insert *Insert
}

func (this *InsertSheet) Model(model interface{}) *InsertSheet {
	this.insert = (&Insert{}).Model(model)
	return this
}

func (this *InsertSheet) initOffset(row *xlsx.Row) (map[string]int, error) {
	header := make(map[string]int)
	for i, c := range row.Cells {
		header[c.Value] = i
	}
	desc := make(map[string][]string)
	for _, c := range this.insert.columns {
		desc[c.Name] = c.Desc
	}
	offset := make(map[string]int)
	for k, s := range desc {
		for _, v := range s {
			if o, ok := header[v]; ok {
				offset[k] = o
			} else {
				return nil, errors.New("field not exist")
			}
		}
	}
	return offset, nil
}

func (this *InsertSheet) Stmt() *Insert {
	return this.insert
}

func (this *InsertSheet) Sheet(sheet *xlsx.Sheet, update bool) error {
	maxRow := sheet.MaxRow
	maxCol := sheet.MaxCol
	if maxRow <= 1 || maxCol <= 0 {
		return errors.New("sheet row or col not illegal")
	}
	offset, err := this.initOffset(sheet.Row(0))
	if err != nil {
		return err
	}
	var values []map[string]interface{}
	for i := 1; i < maxRow; i++ {
		value := make(map[string]interface{})
		cells := sheet.Row(i).Cells
		for k, o := range offset {
			value[k] = strings.Trim(strings.Trim(cells[o].Value, "\t"), " ")
		}
		values = append(values, value)
	}
	return this.insert.Values(values, update)
}
