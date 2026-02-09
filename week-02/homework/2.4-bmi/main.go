package main

import "fmt"

// CalculateBMI вычисляет индекс массы тела
// weightKg - вес в килограммах
// heightCm - рост в сантиметрах
func CalculateBMI(weightKg, heightCm float64) float64 {
	// TODO: реализуй функцию
	// Не забудь перевести сантиметры в метры!
	// Формула: ИМТ = вес (кг) / рост (м)²
	return weightKg / ((heightCm / 100) * (heightCm / 100))
}

// InterpretBMI возвращает категорию по значению ИМТ
func InterpretBMI(bmi float64) string {
	// TODO: реализуй функцию
	// < 18.5: "Недостаточный вес"
	// 18.5 - 24.9: "Норма"
	// 25.0 - 29.9: "Избыточный вес"
	// >= 30: "Ожирение"
	var str string
	if bmi < 18.5 {
		str = "Недостаточный вес"
	}
	if bmi >= 18.5 && bmi <= 24.9 {
		str = "Норма"
	}
	if bmi >= 25 && bmi <= 29.9 {
		str = "Избыточный вес"
	}
	if bmi >= 30 {
		str = "Ожирение"
	}

	return str
}

func main() {
	// Тест 1: вес 70 кг, рост 175 см → ИМТ ≈ 22.9 (Норма)
	bmi := CalculateBMI(70, 175)
	fmt.Printf("Вес: 70 кг, Рост: 175 см\n")
	fmt.Printf("ИМТ: %.1f\n", bmi)
	fmt.Printf("Категория: %s\n\n", InterpretBMI(bmi))

	bmi2 := CalculateBMI(64, 190)
	fmt.Printf("Вес: 64 кг, Рост: 190 см\n")
	fmt.Printf("ИМТ: %.1f\n", bmi2)
	fmt.Printf("Категория: %s\n\n", InterpretBMI(bmi2))

	bmi3 := CalculateBMI(74, 170)
	fmt.Printf("Вес: 74 кг, Рост: 170 см\n")
	fmt.Printf("ИМТ: %.1f\n", bmi3)
	fmt.Printf("Категория: %s\n\n", InterpretBMI(bmi3))

	bmi4 := CalculateBMI(90, 167)
	fmt.Printf("Вес: 90 кг, Рост: 167 см\n")
	fmt.Printf("ИМТ: %.1f\n", bmi4)
	fmt.Printf("Категория: %s\n\n", InterpretBMI(bmi4))

	// TODO: Добавь ещё несколько тестов:
	// - Недостаточный вес
	// - Избыточный вес
	// - Ожирение
}
