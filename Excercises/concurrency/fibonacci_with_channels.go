package main

import "fmt"

// calcFib calculates fibonacci terms
func calcFib(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

// fibonacci return slice of fibonacci terms
func fibonacci(size int) []int {
	fs := make([]int, size)
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < size; i++ {
			fs[i] = <-c
		}
		quit <- true
	}()
	calcFib(c, quit)
	return fs

}

func main() {
	fmt.Printf("%v ", fibonacci(55))
}
