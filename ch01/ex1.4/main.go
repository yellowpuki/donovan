// Dup2 выводит текст каждой строки, которая появляется
// во входных данных более одного раза. Программа читает
// стандартный ввод или список именованных файлов
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileLineCount type
type FileLineCount struct {
	counts map[string]map[string]int
}

// NewFileLineCount create a new object of FileLineCount type
func NewFileLineCount() *FileLineCount {
	return &FileLineCount{
		counts: make(map[string]map[string]int),
	}
}

// CountLines method counted repeated lines in files
func (flc *FileLineCount) CountLines(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if _, ok := flc.counts[f.Name()]; !ok {
			flc.counts[f.Name()] = make(map[string]int)
		}
		flc.counts[f.Name()][scanner.Text()]++
	}
	if scanner.Err() != nil {
		fmt.Printf("fLineCount: CountLines: %v", scanner.Err())
	}
}

func main() {
	flc := NewFileLineCount()

	files := make([]string, len(os.Args[1:]))
	copy(files, os.Args[1:])

	if len(files) == 0 {
		flc.CountLines(os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			flc.CountLines(f)
			f.Close()
		}
	}

	var sb strings.Builder

	for file, dup := range flc.counts {

		flag := false
		sb.Reset()

		for line, n := range dup {
			if n > 1 {
				flag = true
				sb.WriteString(fmt.Sprintf("%d\t%q\n", n, line))
			}
		}
		if flag {
			fmt.Printf("%s\n%s\n", file, sb.String())
		}
	}
}
