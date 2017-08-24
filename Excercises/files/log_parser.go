package main

import (
	"logstats"
	"os"
)

var (
	path string
)

func main() {
	lp := log_parser.NewParsersControl(true, os.Args[1:len(os.Args)])
	lp.ToConsole()
}
