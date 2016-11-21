package main

import "fmt"

func main() {
	i := -4 // here we initializer
	for {
		// this is an incrementor; was known as a post from init-condition-post in previous code
		if i%2 == 0 {
			continue // when line 9 is true, continue will stop this iteration of the loop and start over;
			// when line 9 is false, the code will keep moving down to execute the rest
			i++
		}
		fmt.Println(i)
		if i >= 50 {
			break // break means to stop the loop altogether; break out of it
		}
	}
}
