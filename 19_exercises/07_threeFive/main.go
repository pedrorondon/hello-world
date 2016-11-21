package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			a += i
		} else if i%5 == 0 {
			a += i
		}

	}
	fmt.Println(a)
}
