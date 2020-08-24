package model

import (
	"fmt"
	"ims/util"
	"strings"
)

func COALESCE(fields ...interface{}) string {
	values := make([]string, len(fields))
	for i, v := range fields {
		values[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("COALESCE(%s)", strings.Join(values, ","))
}

func MAX(field string) string {
	return fmt.Sprintf("MAX(%s)", field)
}

func SUM(field string) string {
	return fmt.Sprintf("SUM(%s)", field)
}

func AS(field, alias string) string {
	return field + " AS " + alias
}

func IS(field string, value interface{}) string {
	s := util.StrSlice(value, "'")
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return fmt.Sprintf("%s=%s", field, s[0])
	} else {
		return fmt.Sprintf("%s IN (%s)", field, strings.Join(s, ","))
	}
}

func NOT(field string, value interface{}) string {
	var s = util.StrSlice(value, "'")
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return fmt.Sprintf("%s!=%s", field, s[0])
	} else {
		return fmt.Sprintf("%s NOT IN (%s)", field, strings.Join(s, ","))
	}
}

func LIKE(field string, value string) string {
	return fmt.Sprintf("%s LIKE '%%%v%%'", field, value)
}

func REGEXP(col string, value string) string {
	return fmt.Sprintf("%s REGEXP '%s'", col, value)
}

func NOTREGEXP(col string, value string) string {
	return fmt.Sprintf("%s NOT REGEXP '%s'", col, value)
}

func BETWEEN(col string, f1, f2 interface{}) string {
	return fmt.Sprintf("%s BETWEEN '%v' AND '%v'", col, f1, f2)
}

func GT(col string, v interface{}) string {
	return fmt.Sprintf("%s>'%v'", col, v)
}

func GE(col string, v interface{}) string {
	return fmt.Sprintf("%s>='%v'", col, v)
}

func LT(col string, v interface{}) string {
	return fmt.Sprintf("%s<'%v'", col, v)
}

func LE(col string, v interface{}) string {
	return fmt.Sprintf("%s<='%v'", col, v)
}
