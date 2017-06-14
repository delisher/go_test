package main

import "fmt"

var a int

func main() {
	var mp = map[int](func() int){
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Printf("%v", mp)
}
