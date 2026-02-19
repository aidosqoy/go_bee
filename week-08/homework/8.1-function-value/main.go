package main

import "fmt"

// Apply применяет функцию fn к каждому элементу слайса
func Apply(nums []int, fn func(int) int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс
	// Для каждого элемента вызови fn и добавь результат
	s := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		s[i] = fn(nums[i])
	}
	return s
}

// Filter возвращает слайс элементов, для которых predicate возвращает true
func Filter(nums []int, predicate func(int) bool) []int {
	// TODO: реализуй функцию
	// Если predicate(x) == true, добавь x в результат
	var s []int
	for _, v := range nums {
		if predicate(v) {
			s = append(s, v)
		}
	}
	return s
}

// Reduce сворачивает слайс в одно значение
// initial — начальное значение аккумулятора
// fn — функция (аккумулятор, элемент) -> новый аккумулятор
func Reduce(nums []int, initial int, fn func(int, int) int) int {
	// TODO: реализуй функцию
	// acc = initial
	// для каждого x: acc = fn(acc, x)
	// вернуть acc
	acc := initial
	for _, x := range nums {
		acc = fn(acc, x)
	}
	return acc
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("=== Функции высшего порядка ===")
	fmt.Println("Исходный слайс:", nums)

	// Apply
	doubled := Apply(nums, func(x int) int { return x * 2 })
	fmt.Println("Apply (*2):", doubled)

	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Println("Apply (x²):", squared)

	// Filter
	even := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Filter (чётные):", even)

	greaterThan2 := Filter(nums, func(x int) bool { return x > 2 })
	fmt.Println("Filter (>2):", greaterThan2)

	// Reduce
	sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
	fmt.Println("Reduce (сумма):", sum)

	product := Reduce(nums, 1, func(acc, x int) int { return acc * x })
	fmt.Println("Reduce (произведение):", product)
}
