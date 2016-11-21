package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++ // can also be written as mySlice[0] += 1 or mySlice[0] = mySlice[0] + 1
	fmt.Println(mySlice[0])

}
