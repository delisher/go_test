package stack

import "testing"

func TestPopWorkAfterPush(t *testing.T) {
	s := new(Stack)
	s.Push(11)
	pop_res := s.Pop()
	if pop_res != 11 {
		t.Log("Pop must return 11")
		t.Fail()
	}
}

func TestPushPop(t *testing.T) {
	s := new(Stack)

	s.pushTimes(5)
	if s.data != [10]int{0, 1, 2, 3, 4} {
		t.Log("s.data must be equal [0,1,2,3,4,0,0,0,0,0]")
		t.Fail()
	}
	if s.i != 5 {
		t.Log("s.i must be equal 5")
		t.Fail()
	}

	s.pushTimes(10)
	if s.data != [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		t.Log("s.data must be equal [0,1,2,3,4,5,6,7,8,9]")
		t.Fail()
	}
	if s.i != 10 {
		t.Log("s.i must be equal 10")
		t.Fail()
	}

	s.popTimes(5)
	if s.data != [10]int{0, 1, 2, 3, 4} {
		t.Log("s.data must be equal [0,1,2,3,4,0,0,0,0,0]")
		t.Fail()
	}
	if s.i != 5 {
		t.Log("s.i must be equal 5")
		t.Fail()
	}

	s.popTimes(10)
	if s.data != [10]int{} {
		t.Log("s.data must be equal [0,0,0,0,0,0,0,0,0,0]")
		t.Fail()
	}
	if s.i != 0 {
		t.Log("s.i must be equal 0")
		t.Fail()
	}
}

func (s *Stack) pushTimes(n int) {
	size := s.i
	for i := 0; i < n; i++ {
		s.Push(size + i)
	}
}
func (s *Stack) popTimes(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}
