package orm

import (
	"fmt"
	"log"
	"testing"

	"github.com/tealeg/xlsx"
)

func TestKeyMap(t *testing.T) {
	// go test -v utils_test.go utils.go model.go orm.go

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
