package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"one":   "Good morning",
		"two":   "Bonjour",
		"three": "Buenos Dias",
		"four":  "Bonjourno",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "three")
	fmt.Println(myGreeting)
}

// the resulting output can be displayed in any order, not necessarily the order that is put into the code
