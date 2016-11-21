package main

import "fmt"

func wrap() func() int {
	x := 0 // can also be written var x int which is more idiomatic
	return func() int {
		x++
		return x
	}
}
func main() {
	increment := wrap() // func expression
	fmt.Println(increment())
	fmt.Println(increment())
}
