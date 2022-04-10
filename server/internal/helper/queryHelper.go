package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// lấy các trường ra để validate query
func getFields(i interface{}) []string {
    typOfInterface := reflect.TypeOf(i)
    attrs := make(map[string]bool)

	// fmt.Print(typOfInterface)

    if typOfInterface.Kind() != reflect.Struct {
        fmt.Printf("%v type can't have attributes inspected\n", typOfInterface.Kind())
        return nil
    }

	var rs []string
	
    // loop through the struct's fields and set the map
    for i := 0; i < typOfInterface.NumField(); i++ {
        fieldOfInterface := typOfInterface.Field(i)

        if(fieldOfInterface.Type.Kind() == reflect.Struct){

			// loop field of feild 
            for j := 0 ; j < fieldOfInterface.Type.NumField() ; j++{
                p1 := fieldOfInterface.Type.Field(j)
                v1 := reflect.ValueOf(fieldOfInterface.Type).Elem()
                attrs[p1.Name] = v1.CanSet()
				rs = append(rs, p1.Tag.Get("json")) // FIXME: bị dư một sô thuocj tính mở rọng nhưng tạm thế đã
            }
        }
        if !fieldOfInterface.Anonymous {
            v := reflect.ValueOf(fieldOfInterface.Type)
            v = v.Elem()
            attrs[fieldOfInterface.Name] = v.CanSet()

			rs = append(rs, fieldOfInterface.Tag.Get("json"))
        }
    }

    return rs
}

// check xem query có trogn các trường
func stringInSlice(strSlice []string, s string) bool {

	fmt.Println("check string:",s," in slice: ", strSlice)

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
		return nil, errors.New("unknown field in filter query parameter, only support: " + fmt.Sprint(getFields(model)))
	}
	return map[string]interface{}{field: value}, nil
}
