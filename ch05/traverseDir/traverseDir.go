package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	mode := fileInfo.Mode()
	if mode.IsDir() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}
	path := args[1]
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
