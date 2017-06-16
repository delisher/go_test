// https://en.wikipedia.org/wiki/Linked_list
package main

import (
	"errors"
	"fmt"
)

// Node have it's value, links to previoues and next element
type Node struct {
	Value      interface{}
	prev, next *Node
	list       *List
}

// List have links to first and last element
type List struct {
	head, tail *Node
}

// Front return head node
func (l *List) Front() *Node {
	return l.head
}

// Front return tail node
func (l *List) Back() *Node {
	return l.tail
}

// Prev return next node
func (n *Node) Next() *Node {
	if nxt := n.next; n.list != nil {
		return nxt
	}
	return nil
}

// Prev return previous node
func (n *Node) Prev() *Node {
	if prv := n.prev; n.list != nil {
		return prv
	}
	return nil
}

// Init list with first node
func (l *List) Init(v interface{}) *Node {
	node := &Node{Value: v, list: l}
	l.head = node
	l.tail = node
	return node
}

// Insert new node after "at" node
func (l *List) InsertAfter(v interface{}, at *Node) *Node {
	node := &Node{Value: v, list: l}
	if at == nil {
		l.Init(v)
	} else {
		n := at.next
		if n != nil {
			n.prev = node
			node.next = n
		} else {
			l.tail = node
		}
		at.next = node
		node.prev = at
		node.list = l
	}
	return node
}

// Insert new node before "at" node
func (l *List) InsertBefore(v interface{}, at *Node) *Node {
	node := &Node{Value: v, list: l}
	if at == nil {
		l.Init(v)
	} else {
		p := at.prev
		if p != nil {
			p.next = node
			node.prev = p
		} else {
			l.head = node
		}
		at.prev = node
		node.next = at
		node.list = l
	}
	return node
}

// Push new node after list tail
func (l *List) PushTail(v interface{}) *Node {
	return l.InsertAfter(v, l.Front())
}

// Push new node before list head
func (l *List) PushHead(v interface{}) *Node {
	return l.InsertBefore(v, l.Front())
}

var errEmpty = errors.New("List is empty")

// Pop standart function
func (l *List) Pop() (v interface{}, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}

	return v, err
}

// In main only tests
func main() {
	l := new(List)
	l.PushTail(7)
	n0 := l.PushHead(3)
	n1 := l.PushTail(4)
	l.InsertAfter(6, n1)
	l.InsertAfter(5, n1)
	l.InsertBefore(1, n0)
	l.InsertBefore(2, n0)

	for n := l.Front(); n != nil; n = n.Next() {
		// fmt.Printf("ME %v | PREV %v | NEXT %v\n", n.Value.(int))
		fmt.Printf("%v\n", n)
		// fmt.Printf("%v\n", n.Value.(int))
	}

	fmt.Println()

	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("%v\n", v)
	}
}
