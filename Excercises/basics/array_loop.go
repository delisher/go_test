package main

import "fmt"

func main() {
	const size = 10
	var arr [size]int
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	fmt.Printf("%v\n", arr)
}