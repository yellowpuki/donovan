package main

// Boiling выводит температуру кипения воды
import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Температура кипения воды = %g°F или %g°C\n", f, c)
	// Вывод:
	// Температура кипения воды = 212°F или 100°С
}
