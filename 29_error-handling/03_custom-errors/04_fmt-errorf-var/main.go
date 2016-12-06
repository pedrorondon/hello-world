package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrCust := fmt.Errorf("no square root of a negative number: %v", f) // this code that
		// assigns the error to variable ErrCust has to be placed within func Sqrt because you will
		// not have access to variable f if this line of code is placed at the package level
		return 0, ErrCust
	}
	return 42, nil
}
