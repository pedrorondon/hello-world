package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"one":   "Good morning",
		"two":   "Bonjour",
		"three": "Buenos dias",
		"four":  "Bonjourno",
	}
	fmt.Println(myGreeting)

	delete(myGreeting, "two")

	if x, y := myGreeting["two"]; y { // in the comma ok idiom, the x position is the value, and the y position
		// is a bool of either true or false
		fmt.Println("That value does exist.")
		fmt.Println("value: ", x)
		fmt.Println("exists: ", y)
	} else {
		fmt.Println("That value does not exist.")
		fmt.Println("val: ", x)
		fmt.Println("exists: ", y)
	}
	fmt.Println(myGreeting)
}
