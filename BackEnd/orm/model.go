package orm

import (
	"time"
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
	Cust     string     `gorm:"primary_key" desc:"客户代码"`
	Name     string     `gorm:"default:'NULL'" desc:"客户名称"`
	Inst     string     `gorm:"default:'70300'"`
	OpenDate *time.Time `gorm:"type:date;default:'18991231'"`
	Type     string     `gorm:"default:'NULL'"`
}

func (this *Cust) TableName() string {
	return "cust"
}

type LoanAcct struct {
	Acct         string     `gorm:"column:acct;primary_key;default:'00000000000000000'" desc:"贷款账号"`
	Cust         string     `gorm:"column:cust;default:'00000000000'" desc:"客户代码"`
	Contract     string     `gorm:"column:contract;default:'00000000000000000'" desc:"合同编号"`
	Receipt      string     `gorm:"column:receipt;default:'00000000000000000'" desc:"借据号"`
	Product      string     `gorm:"column:product;default:'00000000'" desc:"核心产品号"`
	Business1    string     `gorm:"column:business_1;default:'NULL'" desc:"业务品种1"`
	Business2    string     `gorm:"column:business_2;default:'NULL'" desc:"业务品种2"`
	Business3    string     `gorm:"column:business_3;default:'NULL'" desc:"业务品种3"`
	Business4    string     `gorm:"column:business_4;default:'NULL'" desc:"业务品种4"`
	Investment   string     `gorm:"column:investment_1;default:'NULL'" desc:"贷款投向"`
	Form         string     `gorm:"column:form;default:'NULL'" desc:"贷款形式"`
	Property     string     `gorm:"column:property;default:'NULL'" desc:"贷款性质"`
	OpenDate     *time.Time `gorm:"column:open_date;type:date;default:'18991231'" desc:"贷款起始日"`
	EndDate      *time.Time `gorm:"column:end_date;type:date;default:'18991231'" desc:"贷款终止日"`
	FirstDate    *time.Time `gorm:"column:first_date;type:date;default:'18991231'" desc:"首次放款日期"`
	Amount       float64    `gorm:"column:amount;default:0.00" desc:"借据金额"`
	Rate         float64    `gorm:"column:rate;default:0.00" desc:"执行年利率"`
	Period       string     `gorm:"column:period;default:'NULL'" desc:"期限类型"`
	Guarantee    string     `gorm:"column:guarantee;default:'NULL'" desc:"担保方式"`
	Repayment    string     `gorm:"column:rapayment;default:'NULL'" desc:"还款方式"`
	RepaymentDay string     `gorm:"column:repayment_day;default:20" desc:"还款日"`
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

func (this *LoanAcct) Desc() {

}

type LoanData struct {
	Acct           string     `gorm:"column:acct;primary_key;index" desc:"贷款账号"`
	State          string     `gorm:"column:state;default:'NULL'" desc:"台帐状态"`
	Balance        float64    `gorm:"column:balance;type:decimal(32,2);default:0.00" desc:"借据余额"`
	DebitCapital   float64    `gorm:"column:debit_capital;type:decimal(32,2);default:0.00" desc:"拖欠本金"`
	DebitIntrest   float64    `gorm:"column:debit_intrest;type:decimal(32,2);default:0.00" desc:"欠息"`
	Classification string     `gorm:"column:classification;default:'NULL'" desc:"五级分类"`
	Date           *time.Time `gorm:"column:date;type:date;primary_key;index"`
}

func (this *LoanData) TableName() string {
	return "loan_data_" + this.Date.Format("20060102")[0:4]
}

type LoanCustOnwer struct {
	Cust string `gorm:"primary_key;index"`
	User string
}

type LoanAcctOwner struct {
	Acct string `gorm:"primary_key;index"`
	User string
}

func CreateTable(table interface{}, withDrop bool) {
	if _db == nil {
		panic("_mysql is nil")
	}

	if withDrop {
		_db.Migrator().DropTable(table)
	}

	_db.Migrator().CreateTable(table)
}
