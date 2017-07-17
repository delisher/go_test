package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// list all files in directory recursively
func main() {
	dir := "d:/work/logs/logs"
	files := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil // not a file.  ignore.
		}
		files = append(files, path)
		return nil
	})

	for _, file := range files {
		fmt.Println(file)
	}
}
