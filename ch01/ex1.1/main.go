// ex1.1 Измените программу echo так,
// чтобы она выводила также os.Args[0],
// имя выполняемой команды
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
