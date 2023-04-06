package main

import (
	"fmt"
	"tempconv/tempconv"
)

const (
	BOILING  = "Температура Кипения"
	FREEZING = "Температура замерзания"
)

func main() {
	bOffset := len(BOILING)
	fOffset := len(FREEZING)
	fmt.Printf("|%*s|%*s|\n", bOffset, BOILING, fOffset, FREEZING)
	fmt.Printf("|%*v|%*v|\n", bOffset, tempconv.BoilingC, fOffset, tempconv.FreezingC)
	fmt.Printf("|%*v|%*v|\n", bOffset, tempconv.CToF(tempconv.BoilingC), fOffset, tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("|%*v|%*v|\n", bOffset, tempconv.CToK(tempconv.BoilingC), fOffset, tempconv.CToK(tempconv.FreezingC))
}
