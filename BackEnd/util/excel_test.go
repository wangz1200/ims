package util

import (
	"log"
	"testing"
)

func TestLoadXls(t *testing.T) {
	file := "D:/Desktop/台账科目.xls"
	book, err := LoadXls(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	sheet := book.GetSheet(1)
	row := sheet.Row(2)
	log.Println(sheet.MaxRow, row.LastCol())

}
