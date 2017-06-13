package main

import "fmt"

func main() {
	i := 0
Loop:
	if i < 10 {
		fmt.Println(i)
		i++
		goto Loop
	}
}