//Buffered Channels

//Process message twice ,instead of processing once and storing : better

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//Casts bidirectional channel into a unidirectional channel
	ch := make(chan int, 50) //Internal data store in channel to store 50 integers
	wg.Add(2)
	go func(ch <-chan int) { //Recieve only channel
		for i := range ch { //But there can be infinite no. of element in channel, so cloase channel in line 27, if u comment out 27, deadlock!
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
