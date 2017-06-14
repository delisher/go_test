package main

import "fmt"

func main() {
	fmt.Printf("%v ", fibonacci(10))
}

func fibonacci(size int) []int { // return only complete result
	result := make([]int, size)
	result[0], result[1] = 1, 1
	for i := 2; i < size; i++ {
		result[i] = result[i - 1] + result[i - 2]
	}
	return result
}