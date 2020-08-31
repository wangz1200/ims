package orm

import (
	"time"
)

func createTable(table interface{}, withDrop bool) {
	if _db == nil {
		panic("_mysql is nil")
	}

	if withDrop {
		_db.DropTable(table)
	}

	_db.CreateTable(table)
}

type User struct {
	User     string `gorm:"primary_key"`
	Password string `gorm:"Default:'000000'"`
	Name     string
	Inst     string
	State    string
}

func (this *User) CreateTable(withDrop bool) {
	createTable(this, withDrop)
}

type Cust struct {
	Cust     string     `name:"客户代码" gorm:"primary_key"`
	Name     string     `name:"客户名称" gorm:"default:'NULL'"`
	Inst     string     `gorm:"default:'70300'"`
	OpenDate *time.Time `gorm:"type:date;default:'18991231'"`
	Type     string     `gorm:"default:'NULL'"`
}

func (this *Cust) TableName() string {
	return "cust"
}

type LoanAcct struct {
	Acct         string     `name:"贷款账号" gorm:"primary_key"`
	Cust         string     `name:"客户代码" gorm:"default:'00000000000'"`
	Contract     string     `name:"合同编号" gorm:"default:'00000000000000000'"`
	Receipt      string     `name:"借据号" gorm:"default:'00000000000000000'"`
	Product      string     `name:"核心产品号" gorm:"default:'00000000'"`
	ProductName  string     `name:"产品名称" gorm:"default:'NULL'"`
	BusinessName string     `name:"业务品种名称" gorm:"default:'NULL'"`
	Form         string     `name:"贷款形式" gorm:"default:'NULL'"`
	Property     string     `name:"贷款性质" gorm:"default:'NULL'"`
	OpenDate     *time.Time `name:"贷款起始日" gorm:"type:date;default:'18991231'"`
	EndDate      *time.Time `name:"贷款终止日" gorm:"type:date;default:'18991231'"`
	FirstDate    *time.Time `name:"首次放款日期" gorm:"type:date;default:'18991231'"`
	Amount       float64    `name:"借据金额" gorm:"default: 0.00"`
	Rate         float64    `name:"执行年利率" gorm:"default:0.00"`
	Period       string     `name:"期限类型" gorm:"default:'NULL'"`
	Guarantee    string     `name:"担保方式" gorm:"default:'NULL'"`
	Investment   string     `name:"贷款投向" gorm:"default:'NULL'"`
	Repayment    string     `name:"还款方式" gorm:"default:'NULL'"`
	RepaymentDay string     `name:"还款日" gorm:"default:20"`
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

func (this *LoanAcct) CreateTable(withDrop bool) {
	createTable(this, withDrop)
}

type LoanData struct {
	Acct           string     `name:"贷款账号" gorm:"primary_key;index"`
	State          string     `name:"台账状态" gorm:"default:'NULL'"`
	Balance        float64    `name:"借据余额" gorm:"default:0.00"`
	DebitCapital   float64    `name:"拖欠本金" gorm:"default:0.00"`
	DebitIntrest   float64    `name:"欠息" gorm:"default:0.00"`
	Classification string     `name:"五级分类" gorm:"default:'NULL'"`
	Date           *time.Time `gorm:"type:date;primary_key;index"`
}

func (this *LoanData) TableName() string {
	date := this.Date
	if date == nil {
		now := time.Now()
		date = &now
	}
	return "loan_data_" + date.Format("20060102")[0:4]
}

func (this *LoanData) CreateTable(withDrop bool) {
	createTable(this, withDrop)
}

type LoanCustOnwer struct {
	Cust string `gorm:"primary_key;index"`
	User string
}

type LoanAcctOwner struct {
	Acct string `gorm:"PrimaryKey;index"`
	User string
}
