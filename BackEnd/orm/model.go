package orm

import (
	"time"
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
	ProductName  string     `gorm:"column:product_name;default:'NULL'" desc:"产品名称"`
	BusinessName string     `gorm:"column:business_name;default:'NULL'" desc:"业务品种名称"`
	Form         string     `gorm:"column:form;default:'NULL'" desc:"贷款形式"`
	Property     string     `gorm:"column:property;default:'NULL'" desc:"贷款性质"`
	OpenDate     *time.Time `gorm:"column:open_date;type:date;default:'18991231'" desc:"贷款起始日"`
	EndDate      *time.Time `gorm:"column:end_date;type:date;default:'18991231'" desc:"贷款终止日"`
	FirstDate    *time.Time `gorm:"column:first_date;type:date;default:'18991231'" desc:"首次放款日期"`
	Amount       float64    `gorm:"column:amount;default:0.00" desc:"借据金额"`
	Rate         float64    `gorm:"column:rate;default:0.00" desc:"执行年利率"`
	Period       string     `gorm:"column:period;default:'NULL'" desc:"期限类型"`
	Guarantee    string     `gorm:"column:guarantee;default:'NULL'" desc:"担保方式"`
	Investment   string     `gorm:"column:investment;default:'NULL'" desc:"贷款投向"`
	Repayment    string     `gorm:"column:rapayment;default:'NULL'" desc:"还款方式"`
	RepaymentDay string     `gorm:"column:repayment_day;default:20" desc:"还款日"`
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

type LoanData struct {
	Acct           string     `gorm:"column:acct;primary_key;index" desc:"贷款账号"`
	State          string     `gorm:"column:state;default:'NULL'" desc:"台帐状态"`
	Balance        float64    `gorm:"column:balance;default:0.00" desc:"借据余额"`
	DebitCapital   float64    `gorm:"column:debit_capital;default:0.00" desc:"拖欠本金"`
	DebitIntrest   float64    `gorm:"column:debit_intrest;default:0.00" desc:"欠息"`
	Classification string     `gorm:"column:classification;default:'NULL'" desc:"五级分类"`
	Date           *time.Time `gorm:"column:date;type:date;primary_key;index"`
}

func (this *LoanData) TableName() string {
	date := this.Date
	if date == nil {
		now := time.Now()
		date = &now
	}
	return "loan_data_" + date.Format("20060102")[0:4]
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
