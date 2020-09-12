package orm

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
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
	Primary bool
}

func parseDate(v string) (*time.Time, error) {
	tt, err := time.Parse("20060102", v)
	if err != nil {
		return nil, err
	} else {
		return &tt, nil
	}
}

func convertFloat(num []uint8) float64 {
	v := string(num)
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func convertDate(d string) string {
	return strings.ReplaceAll(d, "-", "")[0:8]
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

func (this *Insert) Columns() []*Column {
	return this.columns
}

func (this *Insert) Callback(callback ...func(map[string]interface{})) *Insert {
	this.callback = callback
	return this
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

func LoanSheetField() map[string][]string {
	return map[string][]string{
		"Acct":           {"贷款账号"},
		"Cust":           {"客户代码"},
		"Contract":       {"合同编号"},
		"Receipt":        {"借据号"},
		"Product":        {"核心产品号"},
		"Business":       {"业务品种名称"},
		"Investment":     {"贷款投向"},
		"Form":           {"贷款形式"},
		"Property":       {"贷款性质"},
		"OpenDate":       {"贷款起始日"},
		"EndDate":        {"贷款终止日"},
		"FirstDate":      {"首次放款日期"},
		"Amount":         {"借据金额"},
		"Rate":           {"执行年利率"},
		"Period":         {"期限类型"},
		"Guarantee":      {"担保方式"},
		"Repayment":      {"还款方式"},
		"RepaymentDay":   {"还款日"},
		"State":          {"台账状态"},
		"Balance":        {"借据余额"},
		"DebitCapital":   {"拖欠本金"},
		"DebitIntrest":   {"欠息"},
		"Classification": {"五级分类"},
		"Date":           {"日期"},
	}
}

func SheetOffset(row *xlsx.Row, fields map[string][]string) map[string]int {
	header := make(map[string]int)
	for i, c := range row.Cells {
		header[c.Value] = i
	}
	offset := make(map[string]int)
	for k, desc := range fields {
		for _, v := range desc {
			if o, ok := header[v]; ok {
				offset[k] = o
			}
		}
	}
	return offset
}

func (this *Insert) Sheet(sheet *xlsx.Sheet, update bool) error {
	maxRow := sheet.MaxRow
	maxCol := sheet.MaxCol
	if maxRow <= 1 || maxCol <= 0 {
		return errors.New("sheet row or col not illegal")
	}
	var row []string
	for _, cell := range sheet.Row(0).Cells {
		row = append(row, cell.Value)
	}
	offset, err := this.SheetFieldOffset(row)
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
	return this.Values(values, update)
}

type SheetValues struct {
	columns  []*Column
	sheet    *xlsx.Sheet
	fields   map[string][]string
	callback []func(map[string]interface{})
}

func (this *SheetValues) Offset() map[string]int {
	row := this.sheet.Row(0)
	header := make(map[string]int)
	for i, c := range row.Cells {
		header[c.Value] = i
	}
	offset := make(map[string]int)
	for k, desc := range this.fields {
		for _, v := range desc {
			if o, ok := header[v]; ok {
				offset[k] = o
			}
		}
	}
	return offset
}

func (this *SheetValues) Defaults() map[string]string {
	defaults := make(map[string]string)
	for _, c := range this.columns {
		defaults[c.Name] = c.Default
	}
	return defaults
}

func (this *SheetValues) Values() []map[string]interface{} {
	offset := this.Offset()
	defaults := this.Defaults()
	var values []map[string]interface{}
	for i := 1; i < this.sheet.MaxRow; i++ {

	}
	return nil
}

func InsertCustSheet(sheet *xlsx.Sheet, update bool) error {
	insert := (&Insert{}).Model(&Cust{}).
		Callback(func(row map[string]interface{}) {
			if v, ok := row["Cust"]; ok {
				cust := v.(string)
				l := len(cust)
				if l > 11 {
					row["Cust"] = cust[l-11:]
				}
			}
		})
	return insert.Sheet(sheet, update)
}

func InsertLoaAcctSheet(sheet *xlsx.Sheet, update bool) error {
	insert := (&Insert{}).Model(&LoanAcct{}).
		Callback(func(row map[string]interface{}) {
			v := row["Rate"].(string)
			row["Rate"] = strings.ReplaceAll(v, "%", "")
		})
	return insert.Sheet(sheet, update)
}

func InsertLoanDataSheet(date string, sheet *xlsx.Sheet, update bool) error {
	d, err := parseDate(date)
	if err != nil {
		return err
	}
	model := &LoanData{
		Date: d,
	}
	db := DB().Migrator()
	fmt.Println(db.HasTable(model))
	if db.HasTable(model.TableName()) == false {
		if err := db.CreateTable(model); err != nil {
			return err
		}
	}
	insert := (&Insert{}).Model(model).
		Callback(func(row map[string]interface{}) {
			row["Date"] = date
		})

	return insert.Sheet(sheet, update)
}

func LoanTable(date string) (*gorm.DB, error) {
	d, err := parseDate(date)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	loan := DB().Table((&LoanAcct{}).TableName()+" AS acct").
		Joins(fmt.Sprintf("JOIN %s AS data ON %s=%s AND date=?", (&LoanData{Date: d}).TableName(), "acct.acct", "data.acct"), date)
	return loan, nil
}
