package db

import (
	"fmt"
	"reflect"
	"strings"
)

// SliceToBytes slice to bytes
func SliceToBytes(s interface{}) []byte {
	var ss []string

	switch reflect.TypeOf(s).Kind() {
	case reflect.Array, reflect.Slice:
		arrayValue := reflect.ValueOf(s)
		arrayLen := arrayValue.Len()
		if arrayLen > 0 {
			switch arrayValue.Index(0).Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint32, reflect.Uint64:
				ss = make([]string, arrayLen)
				for i := 0; i < arrayLen; i++ {
					ss[i] = fmt.Sprintf("%v", arrayValue.Index(i).Interface())
				}
			case reflect.String:
				ss = make([]string, arrayLen)
				for i := 0; i < arrayLen; i++ {
					ss[i] = fmt.Sprintf("'%v'", arrayValue.Index(i).Interface())
				}
			}
		}
	}

	return []byte(strings.Join(ss, ","))
}
