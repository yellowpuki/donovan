// Изменитье программу fetch так, чтобы к каждому аргументу URL
// автоматически добавлялся префикс http:// в случае отсутствия
// в нем такового
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Println(resp.Status)
	}
}
