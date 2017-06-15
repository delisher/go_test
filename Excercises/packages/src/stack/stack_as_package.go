package stack

import "strconv"

const (
	size = 10
)

// Stack type holds items
type Stack struct {
	i    int
	data [size]int
}

// Push pushes an item on the stack
func (s *Stack) Push(v int) { // map reserved keyword in go
	if s.i > size-1 {
		return
	}
	s.data[s.i] = v
	s.i++
}

// Pop pops an item from the stack
func (s *Stack) Pop() int { // map reserved keyword in go
	if s.i < 1 {
		return 0
	}
	s.i--
	deleted := s.data[s.i]
	s.data[s.i] = 0
	return deleted
}

// String return stack structure for printf
func (s Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
