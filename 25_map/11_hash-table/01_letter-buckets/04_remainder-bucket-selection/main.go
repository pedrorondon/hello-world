package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12) // string() turns the number in paranes into a string
		// dividing the decimal value, i, by 12 and finding the remainder begins to create buckets for all
		// values that have similar remainders when divided by 12
	}
}
