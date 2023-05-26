//Unidirectional channels, polymorphic behaviour

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//casts bidirectional channel into a unidirectional channel
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { //recieve only channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
