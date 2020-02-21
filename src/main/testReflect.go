package main

import (
	"fmt"
	"reflect"
)

// ParseInterface 解析接口
func ParseInterface(data interface{}) {
	v := reflect.ValueOf(data).Elem()
	fmt.Printf("numField: %+v\n", v.NumField())
	t := reflect.TypeOf(data).Elem()
	typeOfType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("%d. %s %s = %v %+v\n", i, typeOfType.Field(i).Name, field.Type(), field.Interface(), t.Field(i).Tag.Get("column"))
	}
}
