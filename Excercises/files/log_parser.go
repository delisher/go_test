package main

import (
	"fmt"
	"log"
	"logstats"
	"os"
)

func main() {
	// ErrorNameRE := regexp.MustCompile(`(?:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2},\d{3}\s).{10,30}`)
	if file, err := os.Open("D:/Work/LOGS/LOGS/10.32.5.14_drv_1.log"); err == nil {
		defer file.Close()
		prs := log_parser.NewParser(true)
		prs.ParseLog(file)
		for k, _ := range prs.Errors {
			e := prs.Errors[k]
			fmt.Printf("%v:\n\t%v\n\t%v раз\n\t\n", e.Name, e.FullErr, e.Number)
		}
	} else {
		log.Fatal(err)
	}
}
