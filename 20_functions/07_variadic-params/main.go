package main

import "fmt"

func main() {
	n := avg(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func avg(sf ...float64) float64 {
	fmt.Println(sf)         // sf is slice which is just a list
	fmt.Printf("%T \n", sf) // %T gives you type like int, float64, string, etc.
	var total float64       // can also be written as total := 0.0
	// total is used to accumulate values
	for _, v := range sf { // range is used to loop over a collection of items
		total += v // can also be written as total = total + v
	}
	return total / float64(len(sf))
}
