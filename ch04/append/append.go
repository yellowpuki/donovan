// appendInt добавляет новый элемент типа int в слайс []int
package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Имеется место для роста. Расширяем слайс
		z = x[:zlen]
	} else {
		// Места для роста нет. Выделяем новый массив.
		// Увеличиваем в два раза для линейной амортизированной сложности.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Встроенная функция
	}

	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%3d\t%v\n", i, cap(y), y)
		x = y
	}
}
