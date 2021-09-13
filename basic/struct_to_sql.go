package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

func createQuery(i interface{}) {
	value := reflect.ValueOf(i)
	t := value.Type()
	if t.Kind() != reflect.Struct {
		fmt.Println("unsupported type")
		return
	}
	table := t.Name()
	data := make([]string, 0, value.NumField())
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		switch field.Kind() {
		case reflect.String:
			data = append(data, fmt.Sprintf("\"%s\"", field.String()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data = append(data, strconv.FormatInt(field.Int(), 10))
		}
	}
	fmt.Printf("insert into %s values(%s)\n", table, strings.Join(data, ", "))
}
