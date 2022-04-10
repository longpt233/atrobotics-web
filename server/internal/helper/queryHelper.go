package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// lấy các trường ra để validate query
func getFields(i interface{}) []string {
    typ := reflect.TypeOf(i)
    attrs := make(map[string]bool)

    if typ.Kind() != reflect.Struct {
        fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
        return nil
    }
	
    // loop through the struct's fields and set the map
    for i := 0; i < typ.NumField(); i++ {
        p := typ.Field(i)

        if(p.Type.Kind() == reflect.Struct){
            for j := 0 ; j < p.Type.NumField() ; j++{
                p1 := p.Type.Field(j)
                v1 := reflect.ValueOf(p.Type).Elem()
                attrs[p1.Name] = v1.CanSet()
            }
        }
        if !p.Anonymous {
            v := reflect.ValueOf(p.Type)
            v = v.Elem()
            attrs[p.Name] = v.CanSet()
        }
    }
    var rs []string
    for key := range attrs{
        rs = append(rs, key)
    }
    return rs
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
	fmt.Println(getFields(model))
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
		return nil, errors.New("unknown field in filter query parameter, only support: " + fmt.Sprint(getFields(model)))
	}
	return map[string]interface{}{field: value}, nil
}
