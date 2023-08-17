package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channel in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	//Recevied Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		//isChannelOpen is true if the channel is not closed
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	//Sent Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 5
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
