package main

import "fmt"

func main() {
	var xyz interface{} = 3
	fmt.Println(xyz + 6) // needs assertion to precisely indicate that xyz is an int
}

// see next example for corrected coding using assertion
