package main

import "fmt"

// IsPrime проверяет, является ли число простым
func IsPrime(n int) bool {
	// TODO: реализуй функцию
	count := 0
	if n >= 2 {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				count++
			}
		}
		return count <= 2
	}
	return false

}

func main() {
	// Тесты
	testCases := []int{-5, 0, 1, 2, 3, 4, 5, 7, 11, 13, 15, 17, 18, 19, 20, 23, 100}

	for _, n := range testCases {
		result := IsPrime(n)
		fmt.Printf("IsPrime(%d) = %v\n", n, result)
	}
}
