//Buffered Channels

//if u send 27 again in the last example, then deadlock
//will this run?

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//casts bidirectional channel into a unidirectional channel
	ch := make(chan int, 50) //Internal data store in channel to store 50 integers
	wg.Add(2)
	go func(ch <-chan int) { //recieve only channel
		i := <-ch
		fmt.Println(i)

		i = <-ch
		fmt.Println(i)

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}

//When to use ? : Sender and reciever work at different frequencies
//Eg : Burst transmission etc.
