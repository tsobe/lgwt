package reflection

import "reflect"

func walk(input interface{}, fn func(name string)) {
	val := getValue(input)

	walkValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		fnResultVal := val.Call(nil)
		for _, resVal := range fnResultVal {
			walkValue(resVal)
		}
	}
}

func getValue(input interface{}) reflect.Value {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
