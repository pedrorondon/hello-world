package main

import "fmt"

var x = 0 // here, x is declared at package level scope
// can also be written as var x int which sets it to 0 value

func increment() int {
	x++
	return x
}
func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
