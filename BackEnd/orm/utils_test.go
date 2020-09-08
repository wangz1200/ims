package orm

import (
	"log"
	"testing"

	"github.com/tealeg/xlsx"
)

// go test -v utils_test.go utils.go model.go orm.go

func TestKeyMap(t *testing.T) {
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
