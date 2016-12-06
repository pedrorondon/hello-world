package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3) // launch function gen and pass in ints 2 and 3

	// This is Fan Out because you have multiple functions, two func sq(), reading/pulling off
	// from a single channel, named variable in
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int { // function gen takes a variatic number of ints and assign that
	// to the variable nums; when you pass in a variatic number of arguments, meaning you have
	// a variatic parameter, you get a slice[]
	fmt.Printf("Type of nums: %T\n", nums) // Just FYI

	out := make(chan int) // create a channel that takes an int and assign taht to variable out
	go func() {
		for _, v := range nums { // range can be used in different ways; here it is being used to
			// range over a slice[] of ints and assign those values to variable v;
			// when you range over a slice or map, you get the index/key and the value;
			// if you just want the key, you write for k := range nums {;
			// if you just want the value, you write for _, v := range nums {; this indicates
			// that you want to throw away the first part, the key
			out <- v // after looping over the slice of ints with line 28 and assigning those
			// values to variable v, here you place those values on channel out
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int { // func sq takes a chan int and returns another channel
	out2 := make(chan int) // create second channel
	go func() {
		for n := range in { // here, range is being used to range or loop over a channel,
			// channel in and assigns any values to variable n
			out2 <- n * n // takes a value, variable n, off of the channel and multiplies
			// it by itself to get its square and puts it on the new channel we created, out2
		}
		close(out2)
	}()
	return out2
}

func merge(cs ...chan int) chan int { // func merge takes a variatic parameter of chan int
	// and that gives us a slice[] of chan int
	fmt.Printf("Type of cs: %T\n", cs) // Just FYI

	out3 := make(chan int) // create a thrird channel
	var wg sync.WaitGroup
	wg.Add(len(cs)) // get length of slice and adding that count to the WaitGroup

	for _, v2 := range cs { // range over the slice of chan int and throw away the key and just
		// get the value, giving you a single chan int during each literation of the loop
		go func(ch chan int) { // this function will take a chan int and assign it to variable ch
			for n := range ch {
				out3 <- n
			}
			wg.Done()
		}(v2)
	}

	go func() {
		wg.Wait()
		close(out3)
	}()
	return out3
}

// Why Fan Out and Fan In?
// So you can have many workers working over multiple cores to process values coming in over
// the assemply line
