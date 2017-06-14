package main

import "fmt"

func main() {
	s := "Hello everybody!\n"
	fmt.Printf(s)
	change_string(s, 's')
}

// // err
// func main() {
// 	var a int
// 	var b int32
// 	b = a + a
// 	b = b + 5
// }

const (
    a = iota
    b
    c string = "0"
)

func change_string(s string, change rune) {
	// s := "hello"
	c := []rune(s)
	c[0] = change
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}