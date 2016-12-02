package main

import "fmt"

func main() {
	var name interface{} = 33
	str, ok := name.(string) // syntax for assertion is .(type asserting it to be)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n") // this line prints since 33 is not a string
	}
}
