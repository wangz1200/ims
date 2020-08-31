package orm

import (
	"errors"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx"
)

func initFieldOffset(table interface{}, sheet *xlsx.Sheet) (map[string]int, error) {
	if sheet.MaxRow < 1 {
		return nil, errors.New("Sheet Rows Not Enough")
	}

	header := sheet.Rows[0]
	fieldMap := make(map[string]int)
	for i := 0; i < sheet.MaxCol; i++ {
		field := header.Cells[i].Value
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

func UpdateFromSheet(table interface{}, sheet *xlsx.Sheet, withUpdate bool) error {
	keyMap, err := initFieldOffset(table, sheet)
	if err != nil {
		return err
	}

	var values []map[string]interface{}
	for i := 1; i < sheet.MaxRow; i++ {
		row := sheet.Rows[i]
		value := make(map[string]interface{})
		for k, o := range keyMap {
			val := row.Cells[o].Value
			val = strings.ReplaceAll(val, " ", "")
			val = strings.ReplaceAll(val, "\t", "")
			value[k] = val
		}
		if v, ok := value["Cust"]; ok {
			val := v.(string)
			l := len(val)
			if l > 11 {
				value["Cust"] = val[l-11:]
			}
		}
		if v, ok := value["Rate"]; ok {
			val := v.(string)
			val = strings.ReplaceAll(val, "%", "")
			value["Rate"] = val
		}
		values = append(values, value)
	}

	if len(values) == 0 {
		return errors.New("data is empty")
	}

	return nil
}

func IntertLoanAcct(values []map[string]interface{}, update bool) {
}
