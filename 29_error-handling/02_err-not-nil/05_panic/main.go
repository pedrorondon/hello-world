package main

import "os"

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("an error happened:", err)
		// log.Println("an error happened:", err)
		// log.Fatalln(err)
		panic(err)
	}
}

// function panic is good is you want to see stack information
// Panic is equivalent to fmt.Println with additional information printed; also gives an
// ext status similar to log.Fatalln

// func panic is found in package builtin from the standard library; this built-in function
// stops normal execution of the current Goroutine

// proper primary error handling is to build functions, packages, api's that when appropriate,
// return errors and when you use those, you check those errors
