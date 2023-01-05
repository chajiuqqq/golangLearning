package main

func main() {
	a := make(chan int)
	b := make(chan string)
	c := make(chan interface{})

	a <- 1
	<-a
	i := <-a
}
