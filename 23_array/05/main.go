package main

import "fmt"

func main() {
	var x [256]string
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = string(i) // taking a number, i, and making it a string that will be stored in a position, x[i]
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v \n,", v, v, []byte(v)) // %v will reveal the letter;
		// []byte(v) reveals what the slice of bytes looks like; UTF8 has a 1 to 4 byte coding scheme

	}
}
