package orm

import (
	"BackEnd/config"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB = nil
)

func Init() error {
	if DB != nil {
		Close()
	}

	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)
	if db, err := gorm.Open("mysql", url); err != nil {
		return err
	} else {
		DB = db
	}

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

func createTable(table interface{}, withDrop bool) {
	if DB == nil {
		panic("DB is nil")
	}

	if withDrop {
		DB.DropTable(table)
	}

	DB.CreateTable(table)
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

type LoanCust struct {
	Cust     string     `alias:"客户代码" gorm:"primary_key"`
	Name     string     `alias:"客户名称"`
	Inst     string     `gorm:"Default:70300"`
	OpenDate *time.Time `gorm:"Default:18991231"`
	Type     string
}

func (this *LoanCust) TableName() string {
	return "cust"
}

func (this *LoanCust) CreateTable(withDrop bool) {
	createTable(this, withDrop)
}

type LoanAcct struct {
	Acct         string     `alias:"贷款账号" gorm:"primary_key"`
	Cust         string     `alias:"客户代码"`
	Contract     string     `alias:"合同编号"`
	Receipt      string     `alias:"借据号"`
	Product      string     `alias:"核心产品号"`
	ProductName  string     `alias:"产品名称"`
	BusinessName string     `alias:"业务品种名称"`
	Form         string     `alias:"贷款形式"`
	Property     string     `alias:"贷款性质"`
	OpenDate     *time.Time `alias:"贷款起始日"`
	EndDate      *time.Time `alias:"贷款终止日"`
	FirstDate    *time.Time `alias:"首次放款日期"`
	Amount       float64    `alias:"借据金额"`
	Rate         float64    `alias:"执行年利率"`
	Period       string     `alias:"客户代码"`
	Guarantee    string     `alias:"担保方式"`
	Investment   string     `alias:"贷款投向"`
	Repayment    string     `alias:"还款方式"`
	RepaymentDay string     `alias:"还款日"`
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

func (this *LoanAcct) CreateTable(withDrop bool) {
	createTable(this, withDrop)
}

type LoanData struct {
	Acct           string     `alias:"贷款账号" gorm:"primary_key;index"`
	State          string     `alias:"台账状态"`
	Balance        float64    `alias:"借据余额"`
	DebitCapital   float64    `alias:"拖欠本金"`
	DebitIntrest   float64    `alias:"欠息"`
	Classification string     `alias:"五级分类"`
	Date           *time.Time `gorm:"primary_key;index"`
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
