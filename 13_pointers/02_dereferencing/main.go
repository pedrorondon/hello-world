package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a, "a")  // 43
	fmt.Println(&a, "b") // 0x20818a220

	var b *int = &a
	fmt.Println(b, "c")  // 0x20818a220
	fmt.Println(*b, "d") // 43
}

// b is an int pointer
// b points to the memory address where an integer is stored
// to see the value in that memory address, add a * in front of b
// this is known as dereferencing
// in this case, the * is an Operator
