package utils

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func StrSlice(value interface{}, quote string) (ret []string) {
	var p = quote + "%v" + quote
	var t = reflect.TypeOf(value)

	switch t.Kind() {
	case reflect.Array:
		if l := t.Len(); l == 0 {
			return
		} else {
			ret = make([]string, l)
			s := reflect.ValueOf(value).Slice(0, l).Interface().([]interface{})
			for i, v := range s {
				ret[i] = fmt.Sprintf(p, v)
			}
		}

	case reflect.Slice:
		s := value.([]interface{})
		if l := len(s); l == 0 {
			return
		} else {
			ret = make([]string, l)
			for i, v := range s {
				ret[i] = fmt.Sprintf(p, v)
			}
		}

	default:
		ret = []string{fmt.Sprintf(p, value)}
	}

	return
}

func ParseDate(v string) (time.Time, error) {
	return time.Parse("20060102", v)
}

func ConvertFloat(num []uint8) float64 {
	v := string(num)
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func ConvertDate(d string) string {
	return strings.ReplaceAll(d, "-", "")[0:8]
}
