package main

import "fmt"

func main() {

	var myMutant = make(map[string]string) // map with both key and value/element of type string;
	// myMutant is the variable declared to a value of a map that has a key and value of type string
	myMutant["Wolverine"] = "is a beast." // assigning a value of type string, "is a beast,"
	// to a key of type string, "Wolverine"; index accessed like slices
	myMutant["Sabertooth"] = "will be defeated." // assigning another value

	fmt.Println(myMutant)
}

// in place of line 6, using var myMutant map[string]string will create a nil reference and there is no append
// function for maps like there are for slices to bring them to life when nil
// using composite literal, var myMutant = map[string]string{}, is another way to creat a map and
// will NOT creat a nil map

// fmt.Println(myMutant == nil) can be used to test if map is nil
