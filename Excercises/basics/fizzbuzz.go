package main

import "fmt"

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	var found bool

	for i := 0; i < 100; i++ {
		found = false
		if i % FIZZ == 0{
			fmt.Printf("Fizz")
			found = true
		}
		if i % BUZZ == 0{
			fmt.Printf("Buzz")
			found = true
		}
		if !found {
			fmt.Printf("%d", i)
			found = true
		}
		fmt.Println()
	}
}