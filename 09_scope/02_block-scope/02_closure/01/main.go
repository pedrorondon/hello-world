package main

import "fmt"

func main() {
	x := 23
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs to the one who is in the ring."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
}
