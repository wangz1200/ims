package orm

import (
	"fmt"
	"ims/config"

	"github.com/jinzhu/gorm"
)

var _mysql *gorm.DB = nil

func MySql() *gorm.DB {
	return _mysql
}

func Connect() error {
	url = fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	if db, err := gorm.Open(pattern, url); err != nil {
		return err
	}

	_mysql = db
	return nil
}

func Close() {
	if _mysql != nil {
		_mysql.Close()
	}
}
