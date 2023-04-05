// Пакет tempconv выполняет вычисления температур
// по Цельсию (Celsius)  и по Фаренгейту(Fahrenheit).
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(C Celsius) Fahrenheit { return Fahrenheit(C*9/5 + 32) }
func FToC(F Fahrenheit) Celsius { return Celsius((F - 32) * 5 / 9) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
