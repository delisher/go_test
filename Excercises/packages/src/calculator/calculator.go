package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var stk = new(stack.Stack)

// https://en.wikipedia.org/wiki/Reverse_Polish_notation
func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				stk.Push(r)
				token = ""
			case c == '*':
				fmt.Printf("%d\n", stk.Pop()*stk.Pop())
			case c == '/':
				r := stk.Pop()
				l := stk.Pop()
				fmt.Printf("%d\n", l/r)
			case c == '+':
				fmt.Printf("%d\n", stk.Pop()+stk.Pop())
			case c == '-':
				r := stk.Pop()
				l := stk.Pop()
				fmt.Printf("%d\n", l-r)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}
