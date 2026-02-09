package main

import "fmt"

// Celsius представляет температуру в градусах Цельсия
type Celsius float64

// Fahrenheit представляет температуру в градусах Фаренгейта
type Fahrenheit float64

// CelsiusToFahrenheit конвертирует Цельсий в Фаренгейт
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	// TODO: реализуй функцию
	// Формула: F = C × 9/5 + 32
	return Fahrenheit(c*9/5 + 32)
}

// FahrenheitToCelsius конвертирует Фаренгейт в Цельсий
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	// TODO: реализуй функцию
	// Формула: C = (F - 32) × 5/9
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	// Тест 1: 100°C = 212°F
	c := Celsius(100)
	f := CelsiusToFahrenheit(c)
	fmt.Printf("%.1f°C = %.1f°F\n", c, f)

	// Тест 2: 32°F = 0°C
	f2 := Fahrenheit(32)
	c2 := FahrenheitToCelsius(f2)
	fmt.Printf("%.1f°F = %.1f°C\n", f2, c2)

	f3 := Fahrenheit(-40)
	c3 := FahrenheitToCelsius(f3)
	fmt.Printf("%.1f°F = %.1f°C\n", f3, c3)

	f4 := Fahrenheit(98.6)
	c4 := FahrenheitToCelsius(f4)
	fmt.Printf("%.1f°F = %.1f°C\n", f4, c4)

	// TODO: Добавь проверку остальных тестовых значений:
	// 0°C = 32°F
	// -40°C = -40°F
	// 37°C = 98.6°F
}
