package main

import "fmt"

// Student представляет информацию о студенте
type Student struct {
	Name  string
	Grade int
}

// GroupByGrade группирует студентов по классу
func GroupByGrade(students []Student) map[int][]Student {
	m := make(map[int][]Student)
	for _, student := range students {
		m[student.Grade] = append(m[student.Grade], student)
	}
	return m
}

// GroupByFirstLetter группирует слова по первой букве
func GroupByFirstLetter(words []string) map[rune][]string {
	// TODO: реализуй функцию
	// Для получения первой буквы используй []rune(word)[0]
	m := make(map[rune][]string)
	for _, word := range words {
		m[[]rune(word)[0]] = append(m[[]rune(word)[0]], word)
	}
	return m
}

// GetStudentNames возвращает имена студентов указанного класса
func GetStudentNames(groups map[int][]Student, grade int) []string {
	// TODO: реализуй функцию
	// Получи слайс студентов по grade
	// Извлеки только имена
	s := make([]string, len(groups[grade]))
	for i, group := range groups[grade] {
		s[i] = group.Name
	}
	return s
}

func main() {
	students := []Student{
		{"Алексей", 10},
		{"Мария", 10},
		{"Иван", 11},
		{"Анна", 11},
		{"Пётр", 10},
	}

	fmt.Println("=== Группировка студентов ===")
	fmt.Println("Все студенты:", students)

	groups := GroupByGrade(students)
	fmt.Println("\nПо классам:")
	for grade, list := range groups {
		fmt.Printf("  %d класс: %v\n", grade, list)
	}

	fmt.Println("\nИмена 10 класса:", GetStudentNames(groups, 10))

	fmt.Println("\n=== Группировка слов ===")
	words := []string{"apple", "apricot", "banana", "avocado", "blueberry"}
	byLetter := GroupByFirstLetter(words)
	fmt.Println("Слова:", words)
	fmt.Println("По первой букве:")
	for k, v := range byLetter {
		fmt.Printf("%c: %v\n", k, v)
	}
}
