package main

import "fmt"

// Greet возвращает строку приветствия для указанного имени
func Greet(name string) string {
	if name == "" {
		name = "Бейтаныс жан"
	}
	return "Сәлем " + name + ", GO әлеміне қош келдің!"
}

func main() {
	message := Greet("Алексей")
	fmt.Println(message)

	fmt.Println(Greet("Мария"))

	// Бонус: проверка пустого имени
	fmt.Println(Greet(""))
}
