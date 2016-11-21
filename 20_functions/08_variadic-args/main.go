package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // data is one item; here it is a slice or list
	// []float64 is of type slice
	n := avg(data...) // ... opens up one item and breaks it up into many pieces that match what
	// a later function is looking for
	fmt.Println(n)
}
func avg(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
