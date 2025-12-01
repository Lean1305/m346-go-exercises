package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) / 1.8
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	f1 := convertCelsiusToFahrenheit(0)
	f2 := convertCelsiusToFahrenheit(25)
	f3 := convertCelsiusToFahrenheit(100)

	fmt.Println(f1, f2, f3)

	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println(convertFahrenheitToCelsius(f1))
	fmt.Println(convertFahrenheitToCelsius(f2))
	fmt.Println(convertFahrenheitToCelsius(f3))
}
