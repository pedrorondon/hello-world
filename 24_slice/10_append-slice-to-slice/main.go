package main

import "fmt"

func main() {
	fslice := []int{1, 2, 3, 4, 5}
	sslice := []int{6, 7, 8, 9}

	fslice = append(fslice, sslice...) // notice the syntax args with the ...

	fmt.Println(fslice)
}
