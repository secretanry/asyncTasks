package main

import "reflect"

func IdentifyType(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	}
	if reflect.TypeOf(val) != nil && reflect.TypeOf(val).Kind() == reflect.Chan {
		return "chan"
	}
	return "unknown"
}
