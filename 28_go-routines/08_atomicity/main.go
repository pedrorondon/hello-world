package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var x sync.WaitGroup
var count int64 // when you see int64 or int32 in concurrency, as opposed to just int,
// it is likely that atomoicity will be used; here, the counter is of type int64
func main() {
	x.Add(2)
	go incrementor("Foo:") // variable incrementor is called twice
	go incrementor("Bar:")
	x.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // sleep for a while;
		atomic.AddInt64(&count, 1)                                 // the 1 increments count by 1
		fmt.Println(s, i, "Counter:", count)
	}
	x.Done()
}

// atomoicity only locks a single thing while a mutex locks a chunk of code;
// creates a variable that can only be accessed by one program thread at a time
