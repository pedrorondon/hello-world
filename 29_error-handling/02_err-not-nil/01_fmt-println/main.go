package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt") // Open function from the os package in the standard
	// library; it returns a pointer to a file, *File, and an error, named variable err;
	// with the underscore_, you are throwing away the pointer to a file because in this code,
	// you are opening a text file that doesn't exist; we just want to catch the error and
	// learn how to deal with it
	if err != nil { // this says if the error is not nil, meaning there is an error, do the
		// following; this is the idiomatic way to deal with errors
		// four different ways you can take the error and do something with that returned error:
		fmt.Println("an error happened:", err) // Println simply prints to standard out which by
		// default, goes to the terminal
		// log.Println("an error happened:", err)
		// log.Fatalln(err)
		// panic(err)
	}
}

// in production code, you should always check for errors and deal with them
