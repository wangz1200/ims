package orm

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB = nil
)

func Init(host, port, user, password, dbname string) error {
	if DB != nil {
		Close()
	}

	pattern := "%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	url := fmt.Sprintf(pattern, user, password, host, port, dbname)
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

type User struct {
	User     string `gorm:"PrimaryKey"`
	Password string `gorm:"Default:'000000'"`
	Name     string
	Inst     string
	State    string
}

type Cust struct {
	Cust     string `gorm:"PrimaryKey"`
	Name     string
	Inst     string     `gorm:"Default:70300"`
	OpenDate *time.Time `gorm:"Default:18991231"`
	Type     string
}

type LoanAcct struct {
	Acct         string     `gorm:"PrimaryKey"`     //账号
	Cust         string     `gorm:"Column:cust_id"` //客户号
	Contract     string     //合同
	Receipt      string     //借据
	Product      string     //产品
	ProductName  string     //产品名称
	Form         string     //贷款形式
	Property     string     //贷款性质
	OpenDate     *time.Time //起始日期
	EndDate      *time.Time //终止日期
	FirstDate    *time.Time //第一次放款日期
	Amount       float64    //借据金额
	Rate         float64    //执行年利率
	Period       string     //贷款期限
	Guarantee    string     //担保方式
	Investment   string     //贷款投向
	Repayment    string     //还款方式
	RepaymentDay string     //还款日
}

func (this *LoanAcct) TableName() string {
	return "loan_acct"
}

type LoanData struct {
	Acct           string     `gorm:"primary_key;index"`
	State          string     //台账状态
	Balance        float64    //借据余额
	DebitCapital   float64    //拖欠本金
	DebitIntrest   float64    //拖欠利息
	Classification string     //五级分类
	Date           *time.Time `gorm:"primary_key;index"`
}

func (this *LoanData) TableName() string {
	date := this.Date
	if date == nil {
		now := time.Now()
		date = &now
	}
	return "dep_data_" + date.Format("20060102")[0:4]
}
