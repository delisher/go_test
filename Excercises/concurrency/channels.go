package main

import "fmt"

// main send iterator to channel
func main() {
	c := make(chan int)
	q := make(chan bool)
	go print_it(c, q)
	for i := 0; i < 10; i++ {
		c <- i
	}
	q <- true
}

// print_it print values from channel c
func print_it(c chan int, q chan bool) {
	for {
		select {
		case i := <-c:
			fmt.Printf("%d\n", i)
		case <-q:
			break
		}
	}
}
