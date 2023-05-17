package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i //Copy of data passes in the channel
		i = 33
		wg.Done()
	}()
	wg.Wait()

}
