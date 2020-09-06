package orm

import (
	"fmt"
	"testing"
)

// go test -v model_test.go model.go orm.go

func TestModel(t *testing.T) {
	Init(true)
	acct := map[string]interface{}{
		"Acct":   "11223344",
		"Amount": "1.33",
	}
	db := DB()
	table := &LoanAcct{}
	err := db.Model(table).Table(table.TableName()).Create(&acct).Error
	fmt.Println(err)
}
