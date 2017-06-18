package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberLinesFlag = flag.Bool("n", false, "number each line")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

// Cat
func cat(r *bufio.Reader) {
	nl := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF && string(buf) == "" {
			break
		}
		if *numberLinesFlag {
			fmt.Fprintf(os.Stdout, "%d %s", nl, buf)
			nl++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
}
