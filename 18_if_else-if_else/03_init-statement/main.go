package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b { // there is an initialization statement in front of the expression for b
		fmt.Println(food)
	}

	if b { // code w/o initialization statement
		fmt.Println("food")
	}
}
