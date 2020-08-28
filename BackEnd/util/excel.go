package util

import (
	_xls "github.com/extrame/xls"
)

type Excel interface {
	Open() error
	Close()

	Book() interface{}

	Sheets(n interface{}) []interface{}

	Cell(int, int) interface{}

	Row(int) []interface{}
	Rows(int, int) [][]interface{}

	Column(int) []interface{}
	Columns(int, int) [][]interface{}

	Range(int, int, int, int) interface{}
}

type _excel struct {
	file string
}

type Xls struct {
	_excel
	book *_xls.WorkBook
}

func (this *Xls) Open(file string) (err error) {
	this.file = file
	this.book, err = _xls.Open(this.file, "utf-8")
	return
}

func (this *Xls) Close() {
	return
}

func (this *Xls) Book() *_xls.WorkBook {
	return this.book
}

func (this *Xls) Cell(row int, col int) interface{} {
	return nil
}

func (this *Xls) Row(row int) []interface{} {
	return nil
}

func (this *Xls) Rows(begin int, end int) [][]interface{} {
	return nil
}

func (this *Xls) Column(col int) []interface{} {
	return nil
}

func (this *Xls) Columns(begin int, end int) interface{} {
	return nil
}

func LoadXls(file string) (*_xls.WorkBook, error) {
	return _xls.Open(file, "utf8")
}
