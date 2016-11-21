package main

import "fmt"

func main() {
	student := make([]string, 35) // 35 is both length and capacity
	pupil := make([][]string, 35)
	student[0] = "Dro" // since a length was given on line 6, append is not needed from position 0 to 34
	fmt.Println(student)
	fmt.Println(pupil)
}
