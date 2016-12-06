package main

import (
	"fmt"
	"log"
	"os"
)

func init() { // func init is a special funtion that initializes things at the startup of
	// your program; it only runs once and you can have as many func init as you like;
	// func init can be in different files and different packages
	//  each source file can define its own niladic init function to set up whatever state
	// is required
	// func init is said to be niladic because it returns nothing; it runs before func main to
	// assist with setup; init is called after all the variable declarations in the package
	//  have evaluated their initializers, and those are evaluated only after all the imported
	// packaged have been initialized
	// func init basically runs before func main to do setup stuff
	nf, err := os.Create("log.txt") // here you are creating a new text file that will catch all
	// of your log messages instead of them being sent to the terminal; when this program is run,
	// nothing appears in the terminal but a new log.txt file appear in the directory;
	// os.Create gives a pointer to a file, named nf, and an error, named err; the pointer
	// to a file implements the Writer inferface
	if err != nil { // if there is an error then proceed to next line of code
		fmt.Println(err) // print the error, named variable err
	}
	log.SetOutput(nf) // sets the output destination for the standard logger to file, log.txt;
	// since SetOutput takes a Writer and we passed in a file which implements the Writer
	// interface, the output desination will be set to this new file, which is pointed to
	// by the pointer to a file, nf

	// Package log implements a simple logging package; it defines a type, Logger, with methods
	// for formatting output; it has a predefined 'standard' Logger accessible through helper
	// functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating
	// a Logger manually.  That logger writes to a stderr and prints the date and time of each
	// logged message
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("an error happened:", err)
		log.Println("an error happened:", err)
		// log.Fatalln(err)
		// panic(err)
	}
}
