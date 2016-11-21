package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus") // fallthtough causes the program to continue executing to next commant;
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian") // since fallthrough is not stated, the program stops running here and does
		// not print the next line
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
