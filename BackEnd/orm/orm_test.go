package orm

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/extrame/xls"
)

func initDB(withDrop bool) {
	cust := &LoanCust{}
	cust.CreateTable(withDrop)

	acct := &LoanAcct{}
	acct.CreateTable(withDrop)

	date, _ := time.Parse("20060102", "20200731")
	data := &LoanData{Date: &date}
	data.CreateTable(withDrop)
}

func TestUpdateLoan(t *testing.T) {
	if err := Init(); err != nil {
		log.Fatal(err)
	} else {
		defer Close()
	}

	book, err := xls.Open("C:/Users/wangz/Desktop/台账科目.xls", "utf-8")
	if err != nil {
		log.Fatal(err)
		return
	}

	sheet := book.GetSheet(1)
	maxRow := sheet.MaxRow

	if maxRow < 0 {
		log.Fatal("sheet empty")
		return
	}

	row := sheet.Row(0)

	fieldMap := make(map[string]int)
	keyMap := make(map[string]int)

	for i := 0; i < row.LastCol(); i++ {
		fieldMap[row.Col(i)] = i
	}

	acct := reflect.TypeOf(LoanAcct{})
	for i := 0; i < acct.NumField(); i++ {
		field := acct.Field(i)
		alias, ok := field.Tag.Lookup("alias")
		if !ok || alias == "" {
			continue
		} else {
			offset, ok := fieldMap[alias]
			if ok {
				keyMap[field.Name] = offset
			} else {
				log.Fatal("not exist")
			}
		}
	}

	log.Println(keyMap)
}
