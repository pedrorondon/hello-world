package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a, "a")  // 43
	fmt.Println(&a, "b") // 0x20818a220

	var b *int = &a
	fmt.Println(b, "c")  // 0x20818a220
	fmt.Println(*b, "d") // 43

	*b = 42             // b says, "The value at this address, change it to 42"
	fmt.Println(a, "e") // 42
}
