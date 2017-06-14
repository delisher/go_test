package main

import "fmt"

func main() {
	s0 := []int{0, -1, 54, 33}
	s1 := append(s0, 345)
	s2 := append(s1, 41, 53, 74)
	s3 := append(s2, s0...)
	s4 := append(s2, -42, -64)

	all_slices := [][]int{s0, s1, s2, s3, s4}
	for _, sl := range all_slices {
		fmt.Printf("%d\n", max(sl))
	}
}

func max(sl []int) (max int) {
	max = sl[0]
	for _, i := range sl {
		if i > max {
			max = i
		}
	}
	return
}
