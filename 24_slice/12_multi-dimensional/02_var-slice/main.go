package main

import "fmt"

func main() {
	var student []string // when not specified, var will always set whatever you are declaring to the 0 value;
	// the 0 zero of a slice is nil because there is no underyling data structure/
	// no address there yet
	var pupil [][]string
	fmt.Println(student)
	fmt.Println(pupil)
	fmt.Println(student == nil) // true since var sets the slice at 0 value

} // like shorhand, this method also requires append
