package main

import "fmt"

// Counter возвращает функцию, которая при каждом вызове
// увеличивает и возвращает счётчик (1, 2, 3, ...)
func Counter() func() int {
	// TODO: реализуй функцию
	// Создай переменную count
	// Верни функцию, которая увеличивает и возвращает count
	i := 0
	return func() int {
		i++
		return i
	}
}

// Accumulator возвращает функцию, которая накапливает сумму
// initial — начальное значение суммы
func Accumulator(initial int) func(int) int {
	// TODO: реализуй функцию
	// Создай переменную sum = initial
	// Верни функцию, которая прибавляет к sum и возвращает результат
	sum := initial
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci возвращает функцию, которая при каждом вызове
// возвращает следующее число Фибоначчи (0, 1, 1, 2, 3, 5, 8, ...)
func Fibonacci() func() int {
	// TODO: реализуй функцию
	// Храни два предыдущих числа
	// При каждом вызове вычисляй следующее
	first := 0
	second := 1
	return func() int {
		result := first
		first, second = second, first+second
		return result
	}
}

func main() {
	fmt.Println("=== Counter ===")
	count := Counter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3

	fmt.Println("\nНезависимый счётчик:")
	count2 := Counter()
	fmt.Println(count2()) // 1

	fmt.Println("\n=== Accumulator ===")
	acc := Accumulator(10)
	fmt.Println("Начало: 10")
	fmt.Println("+5 =", acc(5))
	fmt.Println("+3 =", acc(3))
	fmt.Println("-8 =", acc(-8))

	fmt.Println("\n=== Fibonacci ===")
	fib := Fibonacci()
	fmt.Print("Первые 10 чисел: ")
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()
}
