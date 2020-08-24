package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Map map[string]interface{}
type Row []interface{}
type Rows []Row

var db *sql.DB = nil

func InitMySql(host, port, user, password, name string) (err error) {
	if db != nil {
		_ = db.Close()
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	db, err = sql.Open("mysql", url)
	return
}

func QueryRows(stmt string, args ...interface{}) (ret Rows, err error) {
	if db == nil {
		panic("db is nil")
	}

	var rows *sql.Rows
	rows, err = db.Query(stmt, args ...)
	if err != nil {
		return
	} else {
		defer func() {_ = rows.Close()}()
	}
	colTypes, _ := rows.ColumnTypes()

	for rows.Next() {
		addr := make(Row, len(colTypes))
		for i, t := range colTypes {
			switch t.DatabaseTypeName() {
			case "BIGINT":
				addr[i] = new(int)
			case "DECIMAL":
				addr[i] = new(float64)
			default:
				addr[i] = new(string)
			}
		}
		_ = rows.Scan(addr...)

		val := make(Row, len(colTypes))
		for i, v := range addr {
			switch a := v.(type) {
			case *int:
				val[i] = *a
			case *float64:
				val[i] = *a
			default:
				val[i] = *(a.(*string))
			}
		}
		ret = append(ret, val)
	}

	return
}

func QueryMaps(stmt string, fields []string, args ...interface{}) (ret []Map, err error) {
	if db == nil {
		panic("db is nil")
	}

	var rows *sql.Rows
	rows, err = db.Query(stmt, args ...)
	if err != nil {
		return
	} else {
		defer func() {_ = rows.Close()}()
	}
	colTypes, _ := rows.ColumnTypes()

	if len(colTypes) != len(fields) {
		err = errors.New("fields or columns counts wrong")
		return
	}

	for rows.Next() {
		addr := make(Row, len(colTypes))
		for i, t := range colTypes {
			switch t.DatabaseTypeName() {
			case "BIGINT":
				addr[i] = new(int)
			case "DECIMAL":
				addr[i] = new(float64)
			default:
				addr[i] = new(string)
			}
		}
		_ = rows.Scan(addr...)

		val := make(Map)
		for i, f := range fields {
			switch a := addr[i].(type) {
			case *int:
				val[f] = *a
			case *float64:
				val[f] = *a
			default:
				val[f] = *(a.(*string))
			}
		}

		ret = append(ret, val)
	}

	return
}
