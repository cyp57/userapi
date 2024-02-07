package utils

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// cpnvert struct to primitive.M
func StructToM(v interface{}) (primitive.M, error) {
	result := primitive.M{}
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("bson")
		fmt.Println("i :", i)
		fmt.Println("fieldName :", fieldName)
		fmt.Println("len",len(fieldName))
		if IsEmptyString(fieldName) {
			continue
		}
		if field.Kind() == reflect.Struct && field.Type() == reflect.TypeOf(primitive.ObjectID{}) {
			// If the field is an ObjectID, convert it to Hex
			result[fieldName] = field.Interface().(primitive.ObjectID).Hex()
		} else {
			result[fieldName] = field.Interface()
		}
	}

	return result, nil
}
