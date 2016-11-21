package main

import "fmt"

func main() {
	student := []string{}
	pupil := [][]string{}
	fmt.Println(student)
	fmt.Println(pupil)
	fmt.Println(student == nil) // this asks if student is nil
	// false; slice is not nil since there is an underlying array but append
	// must be used since underlying array does not have any spots in it yet
}
