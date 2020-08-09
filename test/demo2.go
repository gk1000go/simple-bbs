package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

func StructToJsonTagMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Tag.Get("json")
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

func StructToJsonTagMap2(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &result)

	return result
}

type A struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type BB struct {
	ID int
	Name string
}

func main() {
	a := A{1, "keitaj"}
	fmt.Println(a)

	b := StructToMap(&a)
	fmt.Println(b)

	bb := BB{123,"tttt"}
	bx := StructToMap(&bb)
	fmt.Println(bx)

	c := StructToJsonTagMap(&a)
	fmt.Println(c)

	d := StructToJsonTagMap2(&a)
	fmt.Println(d)
}
