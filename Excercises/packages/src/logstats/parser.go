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
	Name          string
	Number        int
	FullErr       string
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
		fmt.Printf("%v errors:\n\n%v\n\n\n%v\n\n\n", e.Number, e.Name, e.FullErr)
	}
}

func (p *Parser) ParseLog() {
	if file, err := os.Open(p.LogFile); err == nil {
		defer file.Close()

		var lastErr bytes.Buffer
		var lastErrText bytes.Buffer

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ln := scanner.Text()
			errName := errTextRE.FindString(ln)
			lastErrName := lastErr.String()
			if ErrorsRE.MatchString(ln) {
				if p.include(lastErrName) && !p.Errors[lastErrName].completeValue {
					p.Errors[lastErrName].FullErr = lastErrText.String()
					p.Errors[lastErrName].completeValue = true
				}
				if p.include(errName) {
					p.Errors[errName].Number++
				} else {
					p.Errors[errName] = &ErrorDescr{errName, 1, "", false}
					lastErrText.Reset()
					lastErrText.WriteString(ln)
					lastErrText.WriteString("\n")
				}
				lastErr.Reset()
				lastErr.WriteString(errName)
				// fmt.Println("=========================", lastErr.String())
			} else {
				if p.include(lastErrName) {
					if !p.Errors[lastErrName].completeValue && !Date_regExp.MatchString(ln) {
						lastErrText.WriteString(ln)
						lastErrText.WriteString("\n")
					} else if Date_regExp.MatchString(ln) {
						p.Errors[lastErrName].FullErr = lastErrText.String()
						p.Errors[lastErrName].completeValue = true
					}
				}
			}
		}
	} else {
		log.Fatal(err)
	}
	return
}
