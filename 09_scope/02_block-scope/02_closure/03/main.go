package main

import "fmt"

func main() {
	x := 2
	increment := func() int { //notice func() with no name in between; this is an anoymous function, a function
		//without a name; func expression: assigning a function to a variable
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

//in comparison to the package below, the one above has a more narrow scope since x is NOT defined at the package
//level

//package main

//import "fmt"

//var x = 0

//func increment() int {
// 	x++
//	return x
//}

//func main() {
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
