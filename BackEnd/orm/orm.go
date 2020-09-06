package orm

import (
	"fmt"
	c "ims/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB = nil

func DB() *gorm.DB {
	return _db
}

func Connect(url string) error {
	Close()

	if db, err := gorm.Open(mysql.Open(url), &gorm.Config{}); err != nil {
		return err
	} else {
		_db = db
	}
	return nil
}

func Close() {
	if _db != nil {
	}
}

func MySqlUrl(host string, port int, user string, password string, name string) string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, name)
}

func Init(drop bool) error {
	url := MySqlUrl(c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)
	if err := Connect(url); err != nil {
		return err
	}

	CreateTable(&Cust{}, drop)
	CreateTable(&LoanAcct{}, drop)
	//CreateTable(&LoanData{}, drop)

	return nil
}
