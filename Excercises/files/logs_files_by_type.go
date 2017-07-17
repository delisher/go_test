package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	filesTypes = map[string]string{
		"drv":  "*drv*",
		"line": "*line*",
		"srv":  "*main*",
	}
	filesByTypes = map[string][]string{}
)

func main() {
	dir := "d:/work/logs"
	filepath.Walk(dir, checkFile)
	for tp, fs := range filesByTypes {
		fmt.Printf("%v:\n", tp)
		for _, f := range fs {
			fmt.Printf("\t%v\n", f)
		}
	}
}

func checkFile(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if f.IsDir() {
		return nil // not a file.  ignore.
	}
	if k, ok := fileType(f.Name()); ok {
		filesByTypes[k] = append(filesByTypes[k], path)
	}
	return nil
}

func fileType(fName string) (string, bool) {
	for tp, exp := range filesTypes {
		matched, err := filepath.Match(exp+".log*", fName)
		if err != nil {
			fmt.Println(err) // malformed pattern
			// return err       // this is fatal.
		}
		if matched {
			return tp, true
		}
	}
	return "", false
}
