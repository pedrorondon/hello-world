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
		return 0, fmt.Errorf("no square root of a negative number: %v", f) // since this
		// is format printing, we use %v, the default value, and give it the float64, f;
		// as a side note, instead of %v, we can be more specific and use %t for a bool,
		// %s for a string, etc.; using %v tells Go to figure out the type for you
		// this allows use to see the value that got passed in along with the customized string
		// string message we created; this gives more context about what happens with this error
		// func Errorf(format string, a ...interface{}) error takes a string and any number
		// of variables
	}
	return 42, nil
}
