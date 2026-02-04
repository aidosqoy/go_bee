package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	x, y := 10, 20
	fmt.Println("До swap:", x, y) // До swap: 10 20

	x, y = swap(x, y)
	fmt.Println("После swap:", x, y) // После swap: 20 10
}
