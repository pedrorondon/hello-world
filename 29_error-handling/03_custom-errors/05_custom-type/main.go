package main

import (
	"fmt"
	"log"
)

type MathError struct { // create a struct and name it MathError
	lat, long string
	err       error // error information is included here
}

func (n *MathError) Error() string { // this function attaches method Error() string
	// to the MathError struct, making the struct of type error; MathError can now be
	// returned as an error
	return fmt.Sprintf("a math error occurred: %v %v %v", n.lat, n.long, n.err) // returns
	// the string in parens() with both text and variables
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		mtherr := fmt.Errorf("no square root of a negative number: %v", f) // create the
		// error with fmt.Errorf() and assign it to variable mtherr
		return 0, &MathError{"50.2289 N", "99.4656 W", mtherr} // return float64 0 and
		// the address to MathError struct with values for variables lat, long, and the
		// error itself, mtherr;
		// this is a technique to include additional specific information along with the
		// error instead of just the error itself
	}

	return 42, nil
}
