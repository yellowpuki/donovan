// Dup1 выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также количество
// ее появлений

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// Примечание: игнорируются потенциальные
	// ошибки из input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
