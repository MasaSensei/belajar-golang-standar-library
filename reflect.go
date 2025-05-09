package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	address string `required:"true" max:"10"`
	city    string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var field reflect.StructField = valueType.Field(i)
		fmt.Println(field.Name, "with type", field.Type)
		fmt.Println(field.Tag.Get("required"))
		fmt.Println(field.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{Name: "Eko"})
	readField(Person{Name: "Eko", address: "Jl. Raya", city: "Bandung"})

	person := Person{Name: "Eko", address: "Jl. Raya", city: "Bandung"}
	fmt.Println(IsValid(person))
}
