package main

import "fmt"

func main() {
	red := fir()   // run function fir which will give a receive only channel and name it red
	f := fact(red) // function fact will take a receive only channel red and give another receive
	// only channel and name it f
	for z := range f { // range over the values f, pull them off of the channel, and name
		// them z
		fmt.Println(z) // print values of variable z
	}
}

func fir() <-chan int { // func fir will return a receive only channel
	out := make(chan int) // create a channel and name it out
	go func() {           // launch a Goroutine
		for i := 0; i < 10; i++ { // loop 10 times from 0-9
			for j := 3; j < 13; j++ { // during each of the 10 loops, loop 10 times from 3-12
				out <- j // from each inner loop, place each of the 10 value on channel out;
				// since that repeats 10 times, 100 values will be placed on the channel
			}
		}
		close(out) // close channel out
	}()
	return out // return the values placed on channel out; will be assigned to variable red on
	// line 6
}

func fact(in <-chan int) <-chan int { // take receive only channel red and name it in
	out2 := make(chan int) // creat a second channel and name it out2
	go func() {
		for q := range in { // range over variable in, a receive only channel, pull the values off
			// of the channel, and name the values q
			out2 <- fact2(q) // call function fact2 for each value of the 100 values, q, and
			// then place those new values on channel out2
		}
		close(out2) // close channel out2
	}()
	return out2 // return the values placed on channel out2;  will be assigned to variable f on
	// line 7
}

func fact2(in2 int) int { // receives ints and names them in2 and returns ints
	total := 1
	for i := in2; i > 0; i-- { // calculates factorial for each of the 100 values passed
		// into the function
		total *= i
	}
	return total
}
