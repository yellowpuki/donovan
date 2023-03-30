// Ex1.2 Измените программу echo так,
// чтобы она выводила индекс и значение каждого аргумента
// по одному аргументу в строке
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, " = ", arg)
	}
}
