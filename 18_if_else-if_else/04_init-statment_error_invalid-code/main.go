package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b { // there is an initialization statement in front of the expression for b
		fmt.Println(food)
	}

	fmt.Println(food) // food is outside of the scope making this line of code invalid

}
