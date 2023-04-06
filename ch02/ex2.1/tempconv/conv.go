package tempconv

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK преобразует температуру по Цельсию в температуру по Кельвину
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF преобразует температуру по Кельвину в температуру по Фаренгейту
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(9*(k-273.15)/5 + 32)
}

// FToK преобразует температуру по Фаренгейту в температуру по кульвину
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(5*(f-32)/9 + 273.15)
}
