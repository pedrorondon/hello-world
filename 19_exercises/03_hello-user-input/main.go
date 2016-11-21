package main

import "fmt"

func main() {
	var a string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&a)
	var b string
	fmt.Scan(&b)
	var c string
	fmt.Scan(&c)
	var d string
	fmt.Scan(&d)
	fmt.Println("Hello", a, b, c, d)

}

// var a-d allows for and requires four names or responses to be entered
