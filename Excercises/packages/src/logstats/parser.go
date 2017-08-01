package log_parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type ErrorDescr struct {
	Name          string
	Number        int
	FullErr       string
	completeValue bool
}

type Parser struct {
	debug  bool
	Errors map[string]*ErrorDescr
}

// type ErrorsInterface interface {
// 	Errors()
// 	allErrorsNumber()
// }

var (
	sensitivity = 50
	errTextRE   = regexp.MustCompile(fmt.Sprintf(`(-\s.{10,%v})`, sensitivity))
	ErrorsRE    = regexp.MustCompile(`.*(\[ERROR\]|\[FATAL\]).+((\n.+){0,2})`)
)

func NewParser(mode bool) *Parser {
	return &Parser{mode, make(map[string]*ErrorDescr)}
}

func (e *Parser) include(name string) bool {
	if _, ok := e.Errors[name]; ok {
		return true
	}
	return false
}

func (p *Parser) ParseLog(logFile io.Reader) {
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		ln := scanner.Text()
		if ErrorsRE.MatchString(ln) {
			str := errTextRE.FindString(ln)
			if p.include(str) {
				p.Errors[str].Number++
			} else {
				p.Errors[str] = &ErrorDescr{str, 1, str, true}
			}
			// fmt.Println(ErrorDescr{str, 1, str, false})
		}
	}
	return
}
