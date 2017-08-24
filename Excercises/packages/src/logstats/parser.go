package log_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

type ErrorDescr struct {
	Name    string
	Number  int
	FullErr *FullErrText
}

type Parser struct {
	debug   bool
	LogFile string
	Errors  map[string]*ErrorDescr
}

var (
	sensitivity = 50
	errTextRE   = regexp.MustCompile(fmt.Sprintf(`(\[ERROR\]|\[FATAL\]).{10,%v}`, sensitivity))
	ErrorsRE    = regexp.MustCompile(`.*(\[ERROR\]|\[FATAL\]).+((\n.+){0,2})`)
	Date_regExp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}`)
)

func (ed *ErrorDescr) complete() bool {
	return ed.FullErr.complete
}

func (ed *ErrorDescr) completeIt() {
	ed.FullErr.complete = true
}

func NewParser(mode bool, path string) *Parser {
	return &Parser{mode, path, make(map[string]*ErrorDescr)}
}

func (p *Parser) logToConsole() {
	for k, _ := range p.Errors {
		e := p.Errors[k]
		fmt.Printf("%v errors:\n\n%v\n\n\n%v\n\n\n", e.Number, e.Name, e.FullErr.String())
	}
}

func isError(str []byte) bool {
	return ErrorsRE.Match(str)
}

func isDated(str []byte) bool {
	return Date_regExp.Match(str)
}

func (p *Parser) addErr(errName string) *ErrorDescr {
	p.Errors[errName] = &ErrorDescr{errName, 1, NewFullError()}
	return p.Errors[errName]
}

func (p *Parser) ParseLog() {
	if file, err := os.Open(p.LogFile); err == nil {
		defer file.Close()

		var lastErr bytes.Buffer
		var errName string
		var lastErrName string

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ln := scanner.Bytes()
			lastErrName = lastErr.String()
			if isError(ln) {
				errName = string(errTextRE.Find(ln))

				if led, ok := p.Errors[lastErrName]; ok {
					if !led.complete() {
						led.completeIt()
					}
				}

				if ed, ok := p.Errors[errName]; ok {
					ed.Number++
				} else {
					ed = p.addErr(errName)
					ed.FullErr.addLine(ln)
				}
				lastErr.Reset()
				lastErr.WriteString(errName)
			} else {
				if led, ok := p.Errors[lastErrName]; ok {
					if !led.complete() && !isDated(ln) {
						led.FullErr.addLine(ln)
					} else if isDated(ln) {
						led.completeIt()
					}
				}
			}
		}
	} else {
		log.Fatal(err)
	}
	return
}
