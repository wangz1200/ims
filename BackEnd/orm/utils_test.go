package orm

import (
	"fmt"
	"log"
	"testing"

	"github.com/tealeg/xlsx"
)

// go test -v utils_test.go utils.go model.go orm.go

func KeyMap(t *testing.T) {
	file := "D:/Desktop/123.xlsx"

	xl, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	sheet := xl.Sheet["原始台账"]
	fmt.Println(sheet.MaxRow)
	if err := UpdateFromSheet(LoanAcct{}, sheet, true); err != nil {
		log.Fatal(err)

	}
}

func TestIntertLoanAcct(t *testing.T) {
	Init(true)
	IntertLoanAcct(nil, true)
}
