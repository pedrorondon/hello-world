package main

import "fmt"

func main() {
	func() { // no receiver, no identifier, no parameteres but could have parameteres
		fmt.Println("I'm studying!!!")
	}()
}
