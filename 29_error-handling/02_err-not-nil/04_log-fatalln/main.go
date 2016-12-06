package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("an error happened:", err)
		// log.Println("an error happened:", err)
		log.Fatalln(err)
		// panic(err)
	}
}

// the Fatal function calls os.Exit(1) after writing the log message
// Fatalln is equivalent to log.Println() followed by a call to os.Exit(1)
// func Exit causes the current program to exit with the given status code; coveentionally,
// code 0 indicates success while non-zero indicates an error; the program then terminates
// immediately
