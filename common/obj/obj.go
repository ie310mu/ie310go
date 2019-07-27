package obj

import (
	"fmt"
	"reflect"
)

//InterfaceIsNil ..
func InterfaceIsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}

	return false
}

//InterfaceToStr ..
func InterfaceToStr(obj interface{}) string {
	msg := fmt.Sprint(obj)
	return msg
}
