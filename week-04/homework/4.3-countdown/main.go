package main

import "fmt"

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
