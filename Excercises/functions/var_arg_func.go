package main

import "fmt"

func main() {
	print_it(1, 2, 3)
	s0 := []int{0, -1, 54, 33}
	print_it()
	print_it(s0...)

}

func print_it(sl ...int) {
	for _, i := range sl {
		fmt.Printf("%d\n", i)
	}
	fmt.Println()
}
