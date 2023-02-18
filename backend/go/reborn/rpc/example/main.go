package main

import (
	"example"
	"fmt"
)

func main() {
	person := &example.Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main Street",
	}
	fmt.Println(person)
}
