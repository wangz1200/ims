package orm

import (
	"fmt"
	"log"
	"testing"

	"github.com/tealeg/xlsx"
)

// go test -v utils_test.go utils.go model.go orm.go

func initDB() {
	file := "D:/Desktop/123.xlsx"
	book, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	sheet := book.Sheets[0]

	Init()
	err = InsertLoanDataSheet("20200831", sheet, true)
	log.Println(err)
}

func TestUtils(t *testing.T) {
	Init()
	var values []map[string]interface{}
	db, err := LoanTable("20200831")
	if err != nil {
		log.Fatal(err)
		return
	}
	db = db.Select("*").Scan(&values)
	fmt.Println(db.Error)
	for _, row := range values {
		fmt.Println(row)
	}
}
