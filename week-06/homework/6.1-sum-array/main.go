package main

import "fmt"

// Sum возвращает сумму всех элементов массива
func Sum(arr [5]int) int {
	// TODO: реализуй функцию
	// Пройди по всем элементам и сложи их
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// Average возвращает среднее арифметическое элементов массива
func Average(arr [5]int) float64 {
	// TODO: реализуй функцию
	// Используй Sum и подели на количество элементов
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum / 5)
}

// Max возвращает максимальный элемент массива
func Max(arr [5]int) int {
	// TODO: реализуй функцию
	// Пройди по массиву и найди максимум
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

// Min возвращает минимальный элемент массива
func Min(arr [5]int) int {
	// TODO: реализуй функцию
	// Пройди по массиву и найди минимум
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}

func main() {
	arr := [5]int{10, 20, 5, 15, 30}

	fmt.Println("=== Операции с массивом ===")
	fmt.Println("Массив:", arr)
	fmt.Println("Сумма:", Sum(arr))
	fmt.Println("Среднее:", Average(arr))
	fmt.Println("Максимум:", Max(arr))
	fmt.Println("Минимум:", Min(arr))
}
