package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

//use waitgroups instead of time.sleep

func main() {
	var msg = "Hi"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	//this last parentheses invokes the function immediately

	// you can use the below to fuck up, pass by reference here due to go providing closure.
	// go func() {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }()

	msg = "Hello"
	wg.Wait()
}
