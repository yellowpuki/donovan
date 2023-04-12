// Dedup отсекает дубликаты на входе
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := map[string]bool{}
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
