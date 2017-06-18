package main

import "fmt"

type e interface{}

func main() {
	s0 := []e{0, -1, 54, 33}
	s1 := append(s0, 41, 53, 74)
	f0 := []e{0.34, -1.3, 54.11, 33.2}
	f1 := append(f0, 41.2, 53.34, 74.2)

	all_slices := [][]e{s0, s1, f0, f1}
	for _, sl := range all_slices {
		fmt.Println(max(sl))
	}
}

func max(sl []e) (max e) {
	max = sl[0]
	for _, i := range sl {
		if Greater(i, max) {
			max = i
		}
	}
	return
}

func Greater(l, r e) bool {
	switch l.(type) {
	case int:
		if ri, ok := r.(int); ok {
			return l.(int) > ri
		}
	case float64:
		ri := r.(float64)
		return l.(float64) > ri

	}
	return false
}
