package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(v interface{}) reflect.Value {
	val := reflect.ValueOf(v)

	// 如果是val是指针，需要先把指针指向的值取出来
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
