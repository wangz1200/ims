package main

import (
	"fmt"
	"ims/orm"
)

func main() {
	if err := orm.Init("127.0.0.1", "3300", "root", "root", "test"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer orm.Close()
	}

	acct := &orm.LoanAcct{}
	orm.DB.DropTable(acct)
	orm.DB.CreateTable(acct)
}
