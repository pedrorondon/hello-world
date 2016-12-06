package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10) //  according to line 17 and 18, it a negative number is passed to
	// func Sqrt, an error will be returned and will be assigned to variable err
	if err != nil { // if there is an error, then..
		log.Fatalln(err) // the error will be logged with a timestamp and followed by a call to
		// os.Exit(1); the error will be sent to standard error which is directed towards
		// the terminal
	}
}

func Sqrt(f float64) (float64, error) { // created a function and named it Sqrt; it takes a
	// float64 and assigns it to variable f; the function returns a float64 and an error
	if f < 0 {
		return 0, errors.New("no square root of a negative number") // if the condition in
		// line 19 is met, 0 will be returned as the float64 and the program will pass
		// back a customized error message
		// errors.New() means that you have a package errors and a function New
		// from Godoc.org, func New takes a string and returns an error;
		// type error is from package builtin and has an interface with method Error()string;
		// anything that has that method will implement the error interface and be type error;
		// func New meets this condition and therefore, returns an error
	}
	// implementation
	return 42, nil
}

// printing out very specific, customized error messages can be very useful for anyone using
// the package you created; provides as much specific information as possible about the error
// that occurred and the context so when people have an error message coming back, they
// can know where it came from and what was happenng when it occurred
