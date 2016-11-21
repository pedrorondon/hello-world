package main

import "fmt"

func main() {
	student := []string{} // slice has no assigned length and array has no assigned capacity
	pupil := [][]string{}
	student[0] = "Dro"
	// student = append(student, "Dro") ; append is needed since a length was not assigned to the slice
	fmt.Println(student)
	fmt.Println(pupil)
}
