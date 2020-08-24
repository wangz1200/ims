package model

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	IdxSelFields = iota
	IdxSelTable
	IdxSelWhere
	IdxSelGroup
	IdxSelOrder
	IdxSelOffset
	IdxSelLimit
	IdxSelMax
)

type TableClause struct {
	clause []string
}

func(this TableClause) join(method string, table string, on ...string) *TableClause {
	clause := method + " " + table
	if len(on) > 0 {
		clause = clause + " ON " + strings.Join(on, " AND ")
	}
	this.clause = append(this.clause, clause)
	return &this
}

func(this TableClause) Join(table string, on ...string) *TableClause {
	return this.join("JOIN", table, on...)
}

func(this TableClause) LeftJoin(table string, on ...string) *TableClause {
	return this.join("LEFT JOIN", table, on...)
}

func(this *TableClause) Str() string {
	return strings.Join(this.clause, " ")
}

type SelectStmt struct {
	db *sql.DB
	clause []string
}

func(this *SelectStmt) Set(idx int, value string) *SelectStmt {
	this.clause[idx] = value
	return this
}

func(this SelectStmt) Fields(fields ...string) *SelectStmt {
	clause := ""
	if len(fields) > 0 {
		clause = strings.Join(fields, ",")
	}
	this.clause[IdxSelFields] = clause
	return &this
}

func(this SelectStmt) From(tables ...interface{}) *SelectStmt {
	clause := ""
	if len(tables) > 0 {
		fields := make([]string, len(tables))
		for i, v := range tables {
			var t string
			switch d := v.(type) {
			case string:
				t = d
			case *TableClause:
				t = d.Str()
			case TableClause:
				t = d.Str()
			}
			fields[i] = t
		}
		clause = "FROM " + strings.Join(fields, ",")
	}
	this.clause[IdxSelTable] = clause
	return &this
}

func(this SelectStmt) Where(filters ...string) *SelectStmt {
	var clause = this.clause[IdxSelWhere]
	if len(filters) == 0 {
		clause = ""
	} else if clause == "" {
		clause = "WHERE " + strings.Join(filters, " AND ")
	} else {
		clause = clause + " AND " + strings.Join(filters, " AND ")
	}
	this.clause[IdxSelWhere] = clause
	return &this
}

func(this SelectStmt) Group(fields ...string) *SelectStmt {
	clause := ""
	if len(fields) > 0 {
		clause = "GROUP BY " + strings.Join(fields, ",")
	}
	this.clause[IdxSelGroup] = clause
	return &this
}

func(this SelectStmt) Order(fields ...string) *SelectStmt {
	clause := ""
	if len(fields) > 0 {
		clause = "ORDER BY " + strings.Join(fields, ",")
	}
	this.clause[IdxSelOrder] = clause
	return &this
}

func(this *SelectStmt) Str() string {
	var str []string
	for _, f := range this.clause {
		if f != "" {
			str = append(str, f)
		}
	}
	return strings.Join(str, " ")
}

func(this *SelectStmt) As(alias string) string {
	return fmt.Sprintf("(%s) AS %s", this.Str(), alias)
}

func(this *SelectStmt) Rows(args ...interface{}) (Rows, error) {
	str := this.Str()
	return QueryRows(str, args...)
}

func(this *SelectStmt) Maps(keys []string, args ...interface{}) ([]Map, error) {
	str := this.Str()
	return QueryMaps(str, keys, args...)
}

func Select() *SelectStmt {
	return &SelectStmt{
		db: db,
	}
}