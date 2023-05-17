//How to use mutexes.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} //to wait for groups of goroutines to complete
var counter = 0
var m = sync.RWMutex{} //to protect data acess

//simple mutex is lock and unlock to ensure one entity uses data,
//RWmutex is many can read,one can write, and if anything is reading , we cant write.

func main() {
	runtime.GOMAXPROCS(100) //basically a tuning variable for threads, pass -1 to see how many
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		//asynchro unlocks the lock
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
