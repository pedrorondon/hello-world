package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var x sync.WaitGroup

func main() {
	x.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	x.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println(s, i, "Counter:", count)
	}
	x.Done()
}
