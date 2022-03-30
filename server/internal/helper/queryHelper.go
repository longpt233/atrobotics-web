package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// lấy các trường ra để validate query
func getFields(model interface{}) []string {
	var field []string
	v := reflect.ValueOf(model)
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}
	return field
}

// check xem query có trogn các trường
func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}
	return false
}

// build câu query sort
func ValidateAndReturnSortQuery(model interface{}, sortBy string) (string, error) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, order := splits[0], splits[1]
	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}
	if !stringInSlice(getFields(model), field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}
	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil
}

// build câu query filter
func ValidateAndReturnFilterMap(model interface{}, filter string) (map[string]interface{}, error) {
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}
	field, value := splits[0], splits[1]
	if !stringInSlice(getFields(model), field) {
		return nil, errors.New("unknown field in filter query parameter")
	}
	return map[string]interface{}{field: value}, nil
}
