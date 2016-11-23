package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}
func HashBucket(x string, y int) int {
	a := int(x[0])
	b := a % y
	return b
}
