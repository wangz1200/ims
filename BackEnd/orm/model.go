package orm

import (
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var (
	LoanField = map[string]string{
		"Acct":           "acct.acct",
		"Cust":           "acct.cust",
		"Contract":       "acct.contract",
		"Product":        "acct.product",
		"ProductName":    "acct.product_name",
		"BusinessName":   "acct.business_name",
		"Form":           "acct.form",
		"Property":       "acct.property",
		"OpenDate":       "acct.open_date",
		"EndDate":        "acct.end_date",
		"FirstDate":      "acct.first_date",
		"Amount":         "acct.amount",
		"Rate":           "acct.rate",
		"Period":         "acct.period",
		"Guarantee":      "acct.guarantee",
		"Investment":     "acct.investment",
		"Repayment":      "acct.repayment",
		"RepaymentDay":   "acct.repayment_day",
		"Balance":        "data.balance",
		"DebitCapital":   "data.debit_capital",
		"DebitIntrest":   "data.debit_intrest",
		"Classification": "data.classification",
		"Date":           "data.date",
	}
	CustSheetFields = map[string][]string{}
	DepSheetFields  = map[string][]string{}
	LoanSheetFields = map[string][]string{
		"Acct":           {"贷款账号"},
		"Cust":           {"客户代码"},
		"Contract":       {"合同编号"},
		"Receipt":        {"借据号"},
		"Product":        {"核心产品号"},
		"Business":       {"业务品种名称"},
		"Investment":     {"贷款投向1"},
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
)

type Table interface {
	TableName() string
}

type User struct {
	User     string `gorm:"primary_key"`
	Password string `gorm:"Default:'000000'"`
	Name     string `gorm:"Default:'NULL'"`
	Inst     string `gorm:"Default:'70300'"`
	State    string `gorm:"Default:'NULL'"`
}

type Cust struct {
	Cust     string    `gorm:"primary_key" desc:"客户代码"`
	Name     string    `gorm:"default:'NULL'" desc:"客户名称"`
	Inst     string    `gorm:"default:'70300'"`
	OpenDate time.Time `gorm:"type:date;default:'18991231'"`
	Type     string    `gorm:"default:'NULL'"`
}

func (this *Cust) TableName() string {
	return "cust"
}

type LoanAcct struct {
	Acct         string    `gorm:"column:acct;primary_key;default:'00000000000000000'"`
	Cust         string    `gorm:"column:cust;default:'00000000000'"`
	Contract     string    `gorm:"column:contract;default:'00000000000000000'"`
	Receipt      string    `gorm:"column:receipt;default:'00000000000000000'"`
	Product      string    `gorm:"column:product;default:'00000000'"`
	Business1    string    `gorm:"column:business_1;default:'NULL'"`
	Business2    string    `gorm:"column:business_2;default:'NULL'"`
	Business3    string    `gorm:"column:business_3;default:'NULL'"`
	Business4    string    `gorm:"column:business_4;default:'NULL'"`
	Investment1  string    `gorm:"column:investment_1;default:'NULL'"`
	Investment2  string    `gorm:"column:investment_2;default:'NULL'"`
	Investment3  string    `gorm:"column:investment_3;default:'NULL'"`
	Investment4  string    `gorm:"column:investment_4;default:'NULL'"`
	Form         string    `gorm:"column:form;default:'NULL'"`
	Property     string    `gorm:"column:property;default:'NULL'"`
	OpenDate     time.Time `gorm:"column:open_date;type:date;default:'18991231'"`
	EndDate      time.Time `gorm:"column:end_date;type:date;default:'18991231'"`
	FirstDate    time.Time `gorm:"column:first_date;type:date;default:'18991231'"`
	Amount       float64   `gorm:"column:amount;default:0.00"`
	Rate         float64   `gorm:"column:rate;default:0.00"`
	Period       string    `gorm:"column:period;default:'NULL'"`
	Guarantee    string    `gorm:"column:guarantee;default:'NULL'"`
	Repayment    string    `gorm:"column:rapayment;default:'NULL'"`
	RepaymentDay string    `gorm:"column:repayment_day;default:20"`
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

type LoanData struct {
	Acct           string    `gorm:"column:acct;primary_key;index" desc:"贷款账号"`
	State          string    `gorm:"column:state;default:'NULL'" desc:"台帐状态"`
	Balance        float64   `gorm:"column:balance;type:decimal(32,2);default:0.00" desc:"借据余额"`
	DebitCapital   float64   `gorm:"column:debit_capital;type:decimal(32,2);default:0.00" desc:"拖欠本金"`
	DebitIntrest   float64   `gorm:"column:debit_intrest;type:decimal(32,2);default:0.00" desc:"欠息"`
	Classification string    `gorm:"column:classification;default:'NULL'" desc:"五级分类"`
	Date           time.Time `gorm:"column:date;type:date;primary_key;index"`
	table          string    `gorm:"-"`
}

func (this *LoanData) TableName() string {
	if this.table == "" {
		return "loan_data"
	} else {
		return this.table
	}
}

type LoanCustOnwer struct {
	Cust string `gorm:"primary_key;index"`
	User string
}

type LoanAcctOwner struct {
	Acct string `gorm:"primary_key;index"`
	User string
}

func CreateTable(table interface{}, withDrop bool) error {
	if _db == nil {
		panic("_mysql is nil")
	}
	if withDrop {
		DB().Migrator().DropTable()
	}
	return DB().Migrator().CreateTable()
}

type Sheet struct {
	keys     []string
	fields   map[string][]string
	callback []func(row map[string]interface{})
}

func (this *Sheet) Keys(keys ...string) {
	this.keys = keys
}

func (this *Sheet) Fields(fields map[string][]string) {
	this.fields = fields
}

func (this *Sheet) Callback(callback ...func(map[string]interface{})) {
	this.callback = callback
}

func (this *Sheet) Values(sheet *xlsx.Sheet) []map[string]interface{} {
	if len(this.fields) == 0 {
		return nil
	}
	header := make(map[string]int)
	for i, c := range sheet.Row(0).Cells {
		header[c.Value] = i
	}
	offset := make(map[string]int)
	for _, k := range this.keys {
		offset[k] = -1
		if field, ok := this.fields[k]; ok {
			for _, f := range field {
				if v, ok := header[f]; ok {
					offset[k] = v
				}
			}
		}
	}
	var values []map[string]interface{}
	for i := 1; i < sheet.MaxRow; i++ {
		cells := sheet.Row(i).Cells
		value := make(map[string]interface{})
		for k, o := range offset {
			if o < 0 {
				continue
			}
			value[k] = cells[o].Value
		}
		for _, f := range this.callback {
			f(value)
		}
		values = append(values, value)
	}
	return values
}

func LoanAcctSheet() *Sheet {
	sheet := &Sheet{}
	sheet.Fields(LoanSheetFields)
	sheet.Keys(
		"Acct",
		"Cust",
		"Contract",
		"Receipt",
		"Product",
		"Business",
		"Investment",
		"Form",
		"Property",
		"OpenDate",
		"EndDate",
		"FirstDate",
		"Amount",
		"Rate",
		"Period",
		"Guarantee",
		"Repayment",
		"RepaymentDay",
	)
	sheet.Callback(func(row map[string]interface{}) {
		for k, v := range row {
			val := v.(string)
			val = strings.Trim(strings.Trim(v.(string), "\t"), " ")
			switch k {
			case "Cust":
				l := len(val)
				if l > 11 {
					val = val[l-11:]
				}
				row[k] = val
			case "Rate":
				row[k] = strings.ReplaceAll(val, "%", "")
			case "Business", "Investment":
				for i, v := range strings.Split(val, "->") {
					row[k+strconv.Itoa(i+1)] = v
				}
				delete(row, k)
			default:
				row[k] = val
			}
		}
	})
	return sheet
}

func LoanDataSheet(date string) *Sheet {
	sheet := &Sheet{}
	sheet.Fields(LoanSheetFields)
	sheet.Keys(
		"Acct",
		"State",
		"Balance",
		"DebitCapital",
		"DebitIntrest",
		"Classification",
		"Date",
	)
	sheet.Callback(func(row map[string]interface{}) {
		for k, v := range row {
			val := v.(string)
			val = strings.Trim(strings.Trim(v.(string), "\t"), " ")
			row[k] = val
		}
		row["Date"] = date
	})
	return sheet
}
