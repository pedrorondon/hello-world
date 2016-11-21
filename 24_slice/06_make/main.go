package main

import "fmt"

func main() {
	mio := make([]int, 0, 5) // 0 is the length and 5 is the capacity of the underlying array

	fmt.Println("--------------------")
	fmt.Println(mio)
	fmt.Println(len(mio))
	fmt.Println(cap(mio))
	fmt.Println("--------------------")
	for i := 0; i < 80; i++ {
		mio = append(mio, i)
		fmt.Println("Value:", mio[i], " Length:", len(mio), "Capacity:", cap(mio))
	}
}

/*
alternate way to make a slice:

package main

import "fmt"

func main() {

mySlice := []int{1,3,5,7,9,11,}
for i, value := range mySlice {
fmt.Println(i, "-", value)
}
}

*/
