package main

import "fmt"

type intstring interface{}

func main() {
	s0 := []intstring{0, -1, 54, 33, 41, 53, 74}
	s1 := []intstring{"a", "b", "c", "z", "y", "x"}

	all_slices := [][]intstring{s0, s1}
	all_functions := []func(intstring) intstring{f0, f1}

	for _, sl := range all_slices {
		for fi, f := range all_functions {
			fmt.Printf("Slice %v Map with function №s%d result is: %v\n", sl, fi, Map(f, sl))
		}
	}
}

func f0(i intstring) intstring {
	sd := "__"
	switch i.(type) {
	case int:
		return i.(int) + i.(int)
	case string:
		return i.(string) + sd
	}
	return i
}

func f1(i intstring) intstring {
	sd := "++++"
	switch i.(type) {
	case int:
		return i.(int) * i.(int)
	case string:
		return i.(string) + sd
	}
	return i
}

func Map(f func(intstring) intstring, sl []intstring) []intstring { // map reserved keyword in go
	r := make([]intstring, len(sl))
	for k, v := range sl {
		r[k] = f(v)
	}
	return r
}
