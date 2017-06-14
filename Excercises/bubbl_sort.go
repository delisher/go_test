package main

import "fmt"

func main() {
	s0 := []int{0, -1, 54, 33}
	s1 := append(s0, 345)
	s2 := append(s1, 41, 53, 74)
	s3 := append(s2, s0...)
	s4 := append(s2, -42, -64)
	empty_slice := []int{}

	all_slices := [][]int{s0, s1, s2, s3, s4, empty_slice}
	for _, sl := range all_slices {
		test_sort(sl)
	}
}

func test_sort(sl []int) {
	fmt.Printf("Before sort: %v\n", sl)
	bubble_sort(sl)
	fmt.Printf("After sort: %v\n\n", sl)
}

func bubble_sort(sl []int) {
	for i := 0; i < len(sl) - 1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
}