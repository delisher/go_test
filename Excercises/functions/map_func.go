package main

import "fmt"

func main() {
	s0 := []int{0, -1, 54, 33}
	s1 := append(s0, 41, 53, 74)

	f0 := func(i int) int {
		return i + i
	}

	f1 := func(i int) int {
		return i * i
	}

	all_slices := [][]int{s0, s1}
	all_functions := []func(int) int{f0, f1}

	for _, sl := range all_slices {
		for fi, f := range all_functions {
			fmt.Printf("Slice %v Map with function №s%d result is: %v\n", sl, fi, Map(f, sl))
		}
	}
}

func Map(f func(int) int, sl []int) []int { // map reserved keyword in go
	r := make([]int, len(sl))
	for k, v := range sl {
		r[k] = f(v)
	}
	return r
}
