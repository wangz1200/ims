package main

import (
	"ims/orm"
	"log"

	"github.com/tealeg/xlsx"
)

func test() {
	orm.Init()
	orm.CreateTable(&orm.LoanAcct{}, false)
	file := "C:/Users/wangz/Desktop/123.xlsx"
	book, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := orm.InsertLoanDataSheet("20200831", book.Sheet["原始台账"], true); err != nil {
		log.Fatal(err)
	}
}

func main() {
	test()
}
