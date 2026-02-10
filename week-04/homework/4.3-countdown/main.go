package main

import "fmt"

// Countdown выводит обратный отсчёт от n до 1, затем "Поехали!"
func Countdown(n int) {
	// TODO: реализуй функцию используя defer
	for i := 1; i <= n; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println("Поехали!")
	fmt.Println("=== Обратный отсчёт ===")
	Countdown(5)
}
