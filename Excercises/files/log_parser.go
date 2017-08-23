package main

import (
	// "fmt"
	"logstats"
	"os"
)

var (
	path string
)

func main() {
	// ErrorNameRE := regexp.MustCompile(`(?:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2},\d{3}\s).{10,30}`)
	// for k, v := range os.Args {
	// 	fmt.Printf("%v => %v\n", k, v)
	// }
	lp := log_parser.NewParsersControl(true, os.Args[1:len(os.Args)])
	lp.ToConsole()
}
