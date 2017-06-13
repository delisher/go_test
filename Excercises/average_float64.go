package main

import "fmt"

func main() {
	s0 := []int{54, 33}
	s1 := append(s0, 345)
	s2 := append(s1, 41, 53, 74)
	s3 := append(s2, s0...)

	all_slices := [][]int{s0, s1, s2, s3}
	for _, arr := range all_slices {
		avg := slice_average(arr)
		fmt.Printf("Average of %v is %.2f\n", arr, avg)
	}
}

func sum_of_arr(arr []int) (sum float64) {

	for _, v := range arr {
		sum += float64(v)
	}
	return
}

func slice_average(arr []int) (avg float64) {
	avg = sum_of_arr(arr) / float64(len(arr))
	return
}