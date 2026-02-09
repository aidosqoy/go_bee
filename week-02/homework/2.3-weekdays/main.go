package main

import "fmt"

// TODO: Определи константы для дней недели с помощью iota
// Понедельник = 1, ..., Воскресенье = 7
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	// TODO: добавь остальные дни
)

// IsWeekend возвращает true, если день является выходным (суббота или воскресенье)
func IsWeekend(day int) bool {
	// TODO: реализуй функцию
	return day == Saturday || day == Sunday
}

func main() {
	fmt.Println("Понедельник:", Monday)
	fmt.Println("Вторник:", Tuesday)
	fmt.Println("Среда:", Wednesday)
	fmt.Println("Четверг:", Thursday)
	fmt.Println("Пятница:", Friday)
	fmt.Println("Суббота:", Saturday)
	fmt.Println("Воскресенье:", Sunday)
	// TODO: выведи остальные дни

	fmt.Println()
	fmt.Println("Понедельник - выходной?", IsWeekend(Monday))
	fmt.Println("Среда - выходной?", IsWeekend(Wednesday))
	fmt.Println("Суббота - выходной?", IsWeekend(Saturday))
	fmt.Println("Воскресенье - выходной?", IsWeekend(Sunday))
	// TODO: проверь субботу и воскресенье
}
