package main

import "fmt"

func main() {
	s0 := []float64{54.23, 33.234}
	s1 := append(s0, 345.345)
	s2 := append(s1, 41.43, 53.1, 74.234)
	s3 := append(s2, s0...)
	empty_slice := []float64{}

	all_slices := [][]float64{s0, s1, s2, s3, empty_slice}
	for _, sl := range all_slices {
		avg := slice_average(sl)
		fmt.Printf("Average of %v is %.2f\n", sl, avg)
	}
}

func sum_of_slice(sl []float64) (sum float64) {
	for _, v := range sl {
		sum += float64(v)
	}
	return
}

func slice_average(sl []float64) (avg float64) {
	slice_length := len(sl)
	switch slice_length {
	case 0:
		avg = 0
	default:
		avg = sum_of_slice(sl) / float64(slice_length)
	}
	return
}