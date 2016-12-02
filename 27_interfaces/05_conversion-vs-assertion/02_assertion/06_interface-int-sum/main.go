package main

import "fmt"

func main() {
	var xyz interface{} = 3
	fmt.Println(xyz.(int) + 6) // .(type) is syntax for assertion; here, you asseted that xyz is
	// an int and can therefore, be added to int 6
}
