package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch { // no expression after switch; when there is no expression, the program will run the first
	// one that is true; otherwise, the default will run
	case len(myFriendsName) == 2: // len is length
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi": // this line is an example of multiple cases
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")

	}

}

// order does matter and the program will ONLY run the first case that is true; if none are true, then
// the program will run the default
