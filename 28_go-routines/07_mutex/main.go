package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var x sync.WaitGroup
var count int      // global counter
var mut sync.Mutex // from sync package, type Mutex

func main() {
	x.Add(2)
	go incrementor("Foo:") // variable incrementor is called twice
	go incrementor("Bar:")
	x.Wait()
	fmt.Println("Final Counter:", count)
}
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond) // sleep for a while;
		// this time delay was moved here to prevent delay within the mutex lock
		mut.Lock() // locks the subsequent code while it is in use by a thread and
		// does not allow other threads to access it until the previous one has completed
		// its task; other threads trying to access or change the variable in the locked
		// code will wait
		x := count // global counter value is stored in x
		x++        // increments the counter by 1
		count = x  // then takes that incremented value and writes it back to the global
		// variable; count++ can replace lines 29 - 31
		fmt.Println(s, i, "Counter:", count)
		mut.Unlock()
	}
	x.Done()
}

// a mutex is a mutual exclusion object; created so that multiple program threads
// can take turns sharing the same resource, such as access to a file
