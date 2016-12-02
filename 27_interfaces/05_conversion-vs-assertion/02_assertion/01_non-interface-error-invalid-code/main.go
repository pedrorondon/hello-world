package main

import "fmt"

func main() {
	name := "Drizzy"
	str, ok := name.(string) // syntax for assertion is .(type asserting it to be)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
