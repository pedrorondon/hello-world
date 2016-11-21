package main

import "fmt"

func main() {
	switch "Medhi" { // switch tells the program what to look for until one of the cases
	// evaluates to true as seen in line 10
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

// no default fallthrough
// fallthrough is optional

/* you can specify fallthrough by explicitly statiing it
   break is not needed like in other languages
*/
