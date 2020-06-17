package tempconv

//Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//Fahrenheit to Kelvin (32°F − 32) × 5/9 + 273.15 = 273.15K
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((Kelvin(f)-32)*5/9 + ZeroK)
}

//Celsius to Kelvin 0°C + 273.15 = 273.15K
func CToK(c Celsius) Kelvin {
	return Kelvin(Kelvin(c) + ZeroK)
}
