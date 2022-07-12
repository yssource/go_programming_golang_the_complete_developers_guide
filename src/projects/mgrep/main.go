package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

/*
get files and directories
change each relative path to an absolute one
read files and show results
go to each directory
do the same
*/

func ls(directory string) ([]string, []string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	subDirs := []string{}
	subFiles := []string{}

	for _, file := range files {
		if file.IsDir() {
			subDirs = append(subDirs, filepath.Join(directory, file.Name()))
		} else {
			subFiles = append(subFiles, filepath.Join(directory, file.Name()))
		}
	}

	return subFiles, subDirs
}

func searchForPattern(wg *sync.WaitGroup, path, pattern string) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		words := strings.Split(line, " ")
		for _, word := range words {
			if strings.ToLower(word) == pattern {
				fmt.Printf("%s:%d %s\n", path, lineNumber, line)
			}
		}
	}
}

func traverse(wg *sync.WaitGroup, pattern, root string) {
	path, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(path)
	if err != nil {
		log.Fatal(err)
	}

	subFiles, subDirs := ls(path)

	for _, file := range subFiles {
		wg.Add(1)
		go searchForPattern(wg, file, pattern)
	}

	for _, dir := range subDirs {
		traverse(wg, pattern, dir)
	}
}

func main() {
	var wg sync.WaitGroup
	args := os.Args[1:]
	if len(args) < 2 {
		panic("please enter the pattern and the starting directory for search")
	}

	pattern := strings.ToLower(args[0])

	startingDir, err := filepath.Abs(args[1])
	if err != nil {
		log.Fatal(err)
	}

	traverse(&wg, pattern, startingDir)

	wg.Wait()
}
