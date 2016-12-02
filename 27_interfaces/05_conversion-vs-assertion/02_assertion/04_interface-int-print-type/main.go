package main

import "fmt"

func main() {
	var xyz interface{} = 7
	fmt.Printf("%T\n", xyz) // no assertion
}
