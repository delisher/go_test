package log_parser

import (
	"fmt"
	// "time"
)

type ParsersControl struct {
	debug   bool
	Parsers []*Parser
}

func NewParsersControl(mode bool, paths []string) *ParsersControl {
	prs := ParsersControl{mode, make([]*Parser, len(paths))}
	for i, path := range paths {
		prs.Parsers[i] = NewParser(true, path)
	}
	return &prs
}

func (pc *ParsersControl) ToConsole() {
	for _, prs := range pc.Parsers {
		// f, e := os.Open(flag.Arg(i))
		// go func() {
		fmt.Println("=====================================================")
		fmt.Printf("%v:\n\n\n", prs.LogFile)
		prs.ParseLog()
		prs.logToConsole()
		fmt.Println("----------------------------------------------------")
		// }()
		// time.Sleep(30 * time.Second)
	}
}
