// Dup2 выводит текст каждой строки, которая появляется
// во входных данных более одного раза. Программа читает
// стандартный ввод или список именованных файлов
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := make([]string, len(os.Args[1:]))
	copy(files, os.Args[1:])

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%q\n", n, line)
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %v", scanner.Err())
	}
}
