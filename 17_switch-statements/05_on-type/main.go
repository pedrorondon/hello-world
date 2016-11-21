package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwithOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")
	}
}
