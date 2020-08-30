package orm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/extrame/xls"
)

func InitKeysOffset(table interface{}, sheet *xls.WorkSheet) (map[string]int, error) {
	if sheet.MaxRow < 1 {
		return nil, errors.New("Sheet Rows Not Enough")
	}

	header := sheet.Row(0)
	fieldMap := make(map[string]int)
	for i := 0; i < int(sheet.MaxRow); i++ {
		field := header.Col(i)
		fieldMap[field] = i
	}

	rt := reflect.TypeOf(table)
	keyMap := make(map[string]int)
	for i := 0; i < int(rt.NumField()); i++ {
		exits := true
		field := rt.Field(i)

		if tag, ok := field.Tag.Lookup("name"); ok {
			for _, n := range strings.Split(tag, "|") {
				if o, ex := fieldMap[n]; ex {
					keyMap[field.Name] = o
					break
				} else {
					exits = false
				}
			}
		}

		if !exits {
			return nil, errors.New("Excel Field Not Exist")
		}
	}

	return keyMap, nil
}

func UpdateFromXls(table interface{}, sheet *xls.WorkSheet, withUpdate bool) error {
	keyMap, err := InitKeysOffset(table, sheet)
	if err != nil {
		return err
	}

	var values []map[string]interface{}
	for i := 1; i < int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		value := make(map[string]interface{})
		for k, o := range keyMap {
			if i == 20 {
				v := row.Col(o)
				fmt.Println(v)
			}
			val := row.Col(o)
			val = strings.ReplaceAll(val, " ", "")
			val = strings.ReplaceAll(val, "\t", "")
			value[k] = val
		}
		values = append(values, value)
	}

	return nil
}
