package main

import "fmt"

func main() {
	rem := 7.25
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var xyz interface{} = 3
	fmt.Printf("%T\n", xyz)
	fmt.Printf("%T\n", int(xyz)) // error: cannot convert xyz of type interface tp type int
	// fmt.Printf("%T\n", xyz.(int)) as seen in this line, must use assertion to specify type
	// for variable xyz
}
