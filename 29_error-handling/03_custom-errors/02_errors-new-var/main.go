package main

import (
	"errors"
	"log"
)

var ErrCust = errors.New("no square root of a negative number") // it is idiomatic
// to have Err prefixed at the beginning of a variable name for a variable that is storing an
// error; having it capitalized makes it accessible outside of the package

// here you are abstractng out a customized error created by errors.New() and assigning
// it to variable ErrCust

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrCust
	}
	return 42, nil
}
