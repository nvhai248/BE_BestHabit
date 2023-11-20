package common

import "reflect"

type Conditions struct {
	Deadline string `json:"deadline" form:"deadline"`
}

func ConvertToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i)
		tag := field.Tag.Get("json")

		if tag != "" && tag != "-" {
			switch value.Kind() {
			case reflect.String:
				if value.Interface() != "" {
					result[tag] = value.Interface()
				}
			case reflect.Int:
				if value.Interface() != 0 {
					result[tag] = value.Interface()
				}
			case reflect.Struct:
				result[tag] = ConvertToMap(value.Interface())
			default:
				result[tag] = value.Interface()
			}
		}
	}

	return result
}
