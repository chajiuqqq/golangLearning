package main

import (
	"fmt"
	"time"
)

var signal = make(chan int)

var outSignal = make(chan bool)

func printInt() {
	for {
		select {

		case a := <-signal:
			fmt.Println(a)
		case <-outSignal:
			fmt.Println("success exit")
			break
		}

	}
}

func main() {

	go printInt()
	for i := 0; i < 10; i++ {
		signal <- i
	}
	outSignal <- true
	time.Sleep(1 * time.Second)
}
