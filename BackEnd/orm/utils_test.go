package orm

import (
	"log"
	"testing"

	"github.com/extrame/xls"
)

func TestKeyMap(t *testing.T) {
	// go test -v utils_test.go utils.go model.go orm.go

	file := "C:/Users/wangz/Desktop/台账科目.xls"

	book, err := xls.Open(file, "utf8")
	if err != nil {
		log.Fatal(err)
		return
	}

	sheet := book.GetSheet(0)
	if err := UpdateFromXls(LoanAcct{}, sheet, true); err != nil {
		log.Fatal(err)
	}
}
