package Forms

import "reflect"

func ifNull(a, b interface{}) interface{} {
	defer func() {
		recover()
	}()
	v := reflect.ValueOf(a)
	if v.IsNil() {
		return b
	}
	return a
}
