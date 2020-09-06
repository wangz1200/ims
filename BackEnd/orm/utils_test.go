package orm

import (
	"log"
	"strings"
	"testing"

	"github.com/tealeg/xlsx"
)

// go test -v utils_test.go utils.go model.go orm.go

func TestKeyMap(t *testing.T) {
	file := "C:/Users/wangz/Desktop/台账科目.xlsx"
	book, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	sheet := book.Sheets[0]
	table := &LoanAcct{}

	Init(true)
	is := (&InsertSheet{}).Model(table)
	is.Stmt().SetCallback(func(value map[string]interface{}) {
		if v, ok := value["Rate"]; ok {
			value["Rate"] = strings.Trim(v.(string), "%")
		}
	})
	if err := is.Sheet(sheet, true); err != nil {
		log.Fatal(err)
	}
}
