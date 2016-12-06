package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("an error happened:", err)
		log.Println("an error happened:", err) // prints  a timestamp
		// log.Fatalln(err)
		// panic(err)
	}
}

// Package log implements a simple logging package; it writes to standard error and
// prints the date and time of each logged message; log.Println calls Output to print
// to the standard logger; arguments are handled in the same manner as fmt.Println but
// with  the additional printing of the date and time/ timestamp

// both standard out/stdout and standard error/stderr usually directed towards the terminal
// unless instructed otherwise; we can tell stderr to print elsewhere if we want

// the Fatal functions call os.Exit(1) after writing the log message

// the Panic functions call panic after writing the log message
