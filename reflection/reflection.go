package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	//	numberOfValues := 0
	//	var getField func(int) reflect.Value
	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		//numberOfValues = val.NumField()
		//getField = val.Field

		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
		//numberOfValues = val.Len()
		//getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			//	Walk(val.MapIndex(key).Interface(), fn)
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFuncResult := val.Call(nil)
		for _, res := range valFuncResult {
			walkValue(res)
		}
	}

	// for i := 0; i < numberOfValues; i++ {
	// 	Walk(getField(i).Interface(), fn)
	// }

	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		Walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.String:
	// 		fn(field.String())
	// 	case reflect.Struct:
	// 		Walk(field.Interface(), fn)
	// 	}
	// }

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
