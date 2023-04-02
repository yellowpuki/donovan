// Изменитье программу fetch так, чтобы к каждому аргументу URL
// автоматически добавлялся префикс http:// в случае отсутствия
// в нем такового
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = fixURL(url, "http://")

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func fixURL(url string, prefix string) string {
	if !strings.HasPrefix(url, prefix) {
		url = prefix + url
	}
	return url
}
