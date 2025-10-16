package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")

	pathSplit := filepath.SplitList(path)

	allResults := make([]string, 0)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// Does this exist?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		// Is it a regular file?
		if !mode.IsRegular() {
			continue
		}

		// Is it executable
		if mode&0111 != 0 {
			if !slices.Contains(allResults, fullPath) {
				allResults = append(allResults, fullPath)
				continue
			}
		}
	}
	for _, k := range allResults {
		fmt.Println(k)
	}
}
