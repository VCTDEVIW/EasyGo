package libs

import (
	."fmt"
	"reflect"
)

func ConvInf2Str(param interface{}) string {
	return param.(string)
}

func ConvInf2Int(param interface{}) int {
	return param.(int)
}

func TypeOfInt(param interface{}) bool {
	if !(reflect.TypeOf(param).Kind() == reflect.Int) {
		Println(Sprintf("Import of {%s} parameter is NOT an integer!", param))
		return false
	} else {
		return true
	}
}

func TypeOfStr(param interface{}) bool {
	if !(reflect.TypeOf(param).Kind() == reflect.String) {
		Println(Sprintf("Import of {%s} parameter is NOT a string!", param))
		return false
	} else {
		return true
	}
}

