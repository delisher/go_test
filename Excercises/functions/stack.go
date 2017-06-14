package main

import (
	"fmt"
	"strconv"
)

const (
	size = 10
)

type stack struct {
	i    int
	data [size]int
}

func main() {
	s := new(stack)
	for i := 0; i < 15; i++ {
		s.push(i)
		fmt.Printf("stack: %v\n", s)
	}

	for i := 0; i < 15; i++ {
		d := s.pop()
		fmt.Printf("stack: %v and deleted %v\n", s, d)
	}
}

func (s *stack) push(v int) { // map reserved keyword in go
	if s.i > size-1 {
		return
	}
	s.data[s.i] = v
	s.i++
}

func (s *stack) pop() int { // map reserved keyword in go
	if s.i < 1 {
		return 0
	}
	s.i--
	fmt.Printf("s.i: %v\n", s.i)
	deleted := s.data[s.i]
	s.data[s.i] = 0
	return deleted
}

func (s stack) String() string {
	var str string
	fmt.Printf("SI: %v\n", s.i)
	for i := 0; i <= s.i-1; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
