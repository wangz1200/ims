package orm

import (
	"log"
	"testing"
	"time"
)

func TestUpdateLoan(t *testing.T) {
	if err := Init(); err != nil {
		log.Fatal(err)
	} else {
		defer Close()
	}

	cust := &LoanCust{}
	acct := &LoanAcct{}

	date, _ := time.Parse("20060102", "20200731")
	data := &LoanData{Date: &date}

	if DB.HasTable(cust) {
		DB.DropTable(cust)
	}
	DB.CreateTable(cust)

	if DB.HasTable(acct) {
		DB.DropTable(acct)
	}
	DB.CreateTable(acct)

	if DB.HasTable(data) {
		DB.DropTable(data)
	}
	DB.CreateTable(&data)
}
