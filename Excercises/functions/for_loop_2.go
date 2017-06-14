package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		print_it(i)
	}
}

func print_it(i int) {
	fmt.Printf("%d\n", i)
}