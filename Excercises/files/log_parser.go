package main

import (
	"fmt"
	"log"
	"logstats"
	"os"
)

func main() {
	// ErrorNameRE := regexp.MustCompile(`(?:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2},\d{3}\s).{10,30}`)
	if file, err := os.Open("C:/LOGS/super_big_log.log"); err == nil {
		defer file.Close()
		prs := log_parser.NewParser(true)
		prs.ParseLog(file)
		for k, _ := range prs.Errors {
			e := prs.Errors[k]
			fmt.Printf("%v errors:\n\n%v\n\n\n", e.Number, e.FullErr)
		}
	} else {
		log.Fatal(err)
	}
}
