package log_parser

import (
	"bufio"
	"bytes"
	"fmt"
	// "io"
	"log"
	"os"
	"regexp"
)

type ErrorDescr struct {
	Name    string
	Number  int
	FullErr *FullErrText
}

type FullErrText struct {
	Text          []byte
	completeValue bool
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

func NewParser(mode bool, path string) *Parser {
	return &Parser{mode, path, make(map[string]*ErrorDescr)}
}

func (e *Parser) include(name string) bool {
	if _, ok := e.Errors[name]; ok {
		return true
	}
	return false
}

func (p *Parser) logToConsole() {
	for k, _ := range p.Errors {
		e := p.Errors[k]
		fmt.Printf("%v errors:\n\n%v\n\n\n%v\n\n\n", e.Number, e.Name, string(e.FullErr.Text))
	}
}

func isError(str []byte) bool {
	return ErrorsRE.Match(str)
}

func NewFullError(str []byte) *FullErrText {
	return &FullErrText{str, false}
}

func (fe *FullErrText) addText(str []byte) {
	// fe.Text = fmt.Sprintf("%v\n%v", fe.Text, str)
	fe.Text = append(fe.Text, str...)
	// fe.Text.WriteString(str)
	// fe.Text.WriteString("\n")
}

func (p *Parser) ParseLog() {
	if file, err := os.Open(p.LogFile); err == nil {
		defer file.Close()

		var lastErr bytes.Buffer

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ln := scanner.Bytes()
			lastErrName := lastErr.String()
			if isError(ln) {
				errName := string(errTextRE.Find(ln))
				if p.include(lastErrName) && !p.Errors[lastErrName].FullErr.completeValue {
					p.Errors[lastErrName].FullErr.completeValue = true
				}
				if p.include(errName) {
					p.Errors[errName].Number++
				} else {
					p.Errors[errName] = &ErrorDescr{errName, 1, NewFullError(ln)}
				}
				lastErr.Reset()
				lastErr.WriteString(errName)
			} else {
				if p.include(lastErrName) {
					if !p.Errors[lastErrName].FullErr.completeValue && !Date_regExp.Match(ln) {
						p.Errors[lastErrName].FullErr.addText(ln)
					} else if Date_regExp.Match(ln) {
						p.Errors[lastErrName].FullErr.completeValue = true
					}
				}
			}
		}
	} else {
		log.Fatal(err)
	}
	return
}
