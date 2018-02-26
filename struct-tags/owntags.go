package main

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

type MyType struct {
	ID   int    `validate:"-"`
	Name string `validate:"string,min=2,max=10"`
	Age  int
}

func main() {

	user := MyType{
		ID:   1,
		Name: "Apan",
	}

	t := reflect.TypeOf(user)

	fmt.Println("Reflect:", t)
	fmt.Println("Name:", t.Name())
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Fields:", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("Field", i, ":", field)
		fmt.Println("   Name:", field.Name)
		fmt.Println("   Type:", field.Type)
		fmt.Println("   Tag:", field.Tag)
		fmt.Println("   GetTag:", field.Tag.Get(tagName))
		getValidatorFromTag(field.Tag.Get(tagName))
	}
}

func getValidatorFromTag(tag string) {
	args := strings.Split(tag, ",")

	for _, part := range strings.Split(tag, ",") {
		switch {
		case part == "string":
			min := args[1]
			max := args[2]
			fmt.Println("      : Got string: Min:", min, "Max:", max)
		}
	}
}
