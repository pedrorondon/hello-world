package main

import "fmt"

func main() {
	x := 13 % 3 // 13 divided by 3 is 4 which would be displayed with x := 13 / 3
	// using the % sign outputs the remainder only
	fmt.Println(x)

	if x == 1 { // note the two = symbols after the x; line 10 - 13 show how to write a
		// condition into the package
		fmt.Println("ODD")

	} else {
		fmt.Println("EVEN")
	}
}
