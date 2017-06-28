// it should print only those items which don't have the same successor
// 'a' 'b' 'a' 'a' 'a' 'c' 'd' 'e' 'f' 'g' => 'a' 'b' 'a' 'c' 'd' 'e' 'f' 'g'
package main

import "fmt"

func main() {
	a := []string{"a", "b", "a", "a", "a", "c", "d", "e", "f", "g"}
	first := ""
	for _, s := range a {
		if first != s {
			fmt.Printf("%v", s)
			first = s
		}
	}
}
