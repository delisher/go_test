package main

import "fmt"

func main() {
	fmt.Println(Panic(panicy))
}
func panicy() {
	var a []int
	// var a [4]int
	a[3] = 5
}

func Panic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("recover() value is: '%v'\n", x)
			b = true
		}
	}()
	f()
	return
}