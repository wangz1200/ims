package orm

import (
	"fmt"
	c "ims/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _db *gorm.DB = nil

func DB() *gorm.DB {
	return _db
}

func Connect(url string) error {
	Close()

	if db, err := gorm.Open("mysql", url); err != nil {
		return err
	} else {
		_db = db
	}
	return nil
}

func Close() {
	if _db != nil {
		_db.Close()
	}
}

func MySqlUrl(host string, port int, user string, password string, name string) string {
	return fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, name)
}

func Init() error {
	url := MySqlUrl(c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)
	if err := Connect(url); err != nil {
		return err
	}

	createTable(&Cust{}, false)
	createTable(&LoanAcct{}, false)
	createTable(&LoanData{}, false)

	return nil
}
