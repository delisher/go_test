package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main read value from os.Stdin and send result to wc func
func main() {
	chars, words, lines := wc(bufio.NewReader(os.Stdin))
	fmt.Fprintf(os.Stdout, "chars:%v\nwords:%v\nlines:%v\n", chars, words, lines)
}

// wc counts number of chars, words and lines
func wc(r *bufio.Reader) (chars, words, lines int) {
	for {
		buf, err := r.ReadString('\n')
		if err != nil {
			return
		} else {
			chars += len(buf)
			words += len(strings.Fields(buf))
			lines++
		}
	}
}
