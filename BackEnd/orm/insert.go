package orm

import (
	"ims/utils"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Column struct {
	Name    string
	DBName  string
	Default string
	Primary bool
}

type Insert struct {
	db       *gorm.DB
	columns  []*Column
	callback []func(map[string]interface{})
}

func (this *Insert) DB() *gorm.DB {
	return this.db
}

func (this *Insert) Model(model interface{}) {
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
}

func (this *Insert) Columns() []*Column {
	return this.columns
}

func (this *Insert) Callback(callback ...func(map[string]interface{})) {
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

func InsertLoanAcctSheet(sheet *xlsx.Sheet, update bool) error {
	insert := &Insert{}
	insert.Model(&LoanAcct{})
	return insert.Values(LoanAcctSheet().Values(sheet), update)
}

func InsertLoanDataSheet(date string, sheet *xlsx.Sheet, update bool) error {
	_, err := utils.ParseDate(date)
	if err != nil {
		return err
	}
	model := LoanData{
		table: "loan_data_" + date[0:4],
	}
	if err := CreateTable(model, false); err != nil {
		return err
	}
	insert := &Insert{}
	insert.Model(model)
	return insert.Values(LoanDataSheet(date).Values(sheet), update)
}
