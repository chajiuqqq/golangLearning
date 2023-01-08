package main

import (
	"fmt"
)

func f() {

}

var s = make(chan int)

func main() {
	go f()
	for i := 0; i < 10; i++ {
		fmt.Println(<-s)
	}
}
