package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}

	nested_kind := reflect.ValueOf(nested).Kind()
	if nested_kind == reflect.Slice {
		newNested := nested.([]interface{})
		for _, elem := range newNested {
			result = append(result, Flatten(elem)...)
		}
	} else if nested_kind == reflect.Int {
		result = append(result, nested)
	}

	return result

}
