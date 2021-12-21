package ds

import (
	"reflect"
)

func inArray(haystack interface{}, needle interface{}) bool {
	arr := reflect.ValueOf(haystack)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == needle {
			return true
		}
	}

	return false
}
