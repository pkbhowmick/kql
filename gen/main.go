package main

import (
	"fmt"
	"reflect"
	"strings"

	core "k8s.io/api/core/v1"
)

const (
	JSONSuffix = ",omitempty"
)

func gen(obj interface{}) {
	objType := obj.(reflect.Type)
	for i := 0; i < objType.NumField(); i++ {
		switch objType.Field(i).Type.Kind() {
		case reflect.Slice:
			if objType.Field(i).Type.Elem().Kind() == reflect.Struct {
				gen(objType.Field(i).Type.Elem())
			} else {
				fmt.Println(objType.Field(i).Name, " ", strings.TrimSuffix(objType.Field(i).Tag.Get("json"), JSONSuffix))
			}
		case reflect.Struct:
			gen(objType.Field(i).Type)
		default:
			fmt.Println(objType.Field(i).Name, " ", strings.TrimSuffix(objType.Field(i).Tag.Get("json"), JSONSuffix))
		}
	}
}

func main() {
	objType := reflect.TypeOf(core.Pod{}.ObjectMeta)
	gen(objType)
}
