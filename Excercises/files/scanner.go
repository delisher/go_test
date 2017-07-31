package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	ErrorsRE := regexp.MustCompile(`.*(\[ERROR\]|\[FATAL\]).+((\n.+){0,2})`)
	// ErrorNameRE := regexp.MustCompile(`(?:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2},\d{3}\s).{10,30}`)
	if file, err := os.Open("D:/Work/LOGS/LOGS/10.32.5.14_drv_1.log"); err == nil {
		defer file.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if ErrorsRE.MatchString(scanner.Text()) {
				fmt.Println(ErrorsRE.FindString(scanner.Text()))
			}
		}

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
}
