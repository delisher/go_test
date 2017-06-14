package main

import "fmt"

func main() {
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
	fmt.Printf("%v\n", plusX(2)(2))
	fmt.Printf("%v\n", plusX(8)(2))
}

func plusTwo() func(int) int {
	return func(i int) int { return i + 2 }
}

func plusX(x int) func(int) int {
	return func(i int) int { return i + x }
}
