package main

import "fmt"

// AppendUnique добавляет элемент в слайс только если его там ещё нет
func AppendUnique(slice []int, value int) []int {
	// TODO: реализуй функцию
	// Проверь, есть ли value в slice
	// Если нет — добавь через append
	u := 0
	for _, v := range slice {
		if v == value {
			u++
		}
	}
	if u == 0 {
		slice = append(slice, value)
		return slice
	}
	return slice
}

// RemoveAt удаляет элемент по индексу
func RemoveAt(slice []int, index int) []int {
	// TODO: реализуй функцию
	// Проверь, что index в допустимых границах
	// Используй append для соединения частей до и после index
	if index > len(slice)-1 {
		return slice
	}
	s := append(slice[:index], slice[index+1:]...)
	return s

}

// RemoveValue удаляет первое вхождение значения из слайса
func RemoveValue(slice []int, value int) []int {
	// TODO: реализуй функцию
	// Найди индекс value и используй RemoveAt
	for i, s := range slice {
		if s == value {
			return RemoveAt(slice, i)
		}
	}
	return nil
}

// SumAll суммирует любое количество чисел (variadic функция)
func SumAll(nums ...int) int {
	// TODO: реализуй функцию
	// nums — это слайс, пройди по нему и сложи
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("=== Операции со слайсами ===")
	fmt.Println("Исходный слайс:", nums)

	// Добавление уникальных
	nums = AppendUnique(nums, 6)
	fmt.Println("После AppendUnique(6):", nums)
	nums = AppendUnique(nums, 3)
	fmt.Println("После AppendUnique(3):", nums)

	// Удаление по индексу
	nums = RemoveAt(nums, 2)
	fmt.Println("После RemoveAt(2):", nums)

	// Удаление по значению
	nums = RemoveValue(nums, 5)
	fmt.Println("После RemoveValue(5):", nums)

	// Variadic функция
	fmt.Println("\n=== Variadic функция ===")
	fmt.Println("SumAll(1, 2, 3):", SumAll(1, 2, 3))
	fmt.Println("SumAll(nums...):", SumAll(nums...))
}
