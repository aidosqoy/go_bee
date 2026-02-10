package main

import "fmt"

func Grade(score int) string {
	// TODO: реализуй с помощью switch без условия
	switch {
	case score < 0 || score > 100:
		return "Некорректный балл"
	case score >= 90:
		return "A"
	case score >= 80 && score <= 89:
		return "B"
	case score >= 70 && score <= 79:
		return "C"
	case score >= 60 && score <= 69:
		return "D"
	default:
		return "F"
	}
}

func main() {
	scores := []int{-5, 0, 45, 59, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105}

	for _, score := range scores {
		fmt.Printf("Балл %d → %s\n", score, Grade(score))
	}
}
