package main

import (
	"fmt"
	"strings"
)

// WordCount возвращает map с количеством каждого слова в тексте
func WordCount(text string) map[string]int {
	// TODO: реализуй функцию
	// Используй strings.Fields для разбиения на слова
	// Пройди по словам и увеличивай счётчик в map

	arr := strings.Fields(text)
	m := make(map[string]int)

	for _, v := range arr {
		m[v]++
	}
	return m

}

// CharCount возвращает map с количеством каждого символа
func CharCount(text string) map[rune]int {
	// TODO: реализуй функцию
	// Пройди по строке с range — получишь rune
	m := make(map[rune]int)
	for _, v := range text {
		if v != ' ' {
			m[v]++
		}
	}
	return m
}

// MostFrequent возвращает самое часто встречающееся слово
func MostFrequent(text string) string {
	// TODO: реализуй функцию
	// Используй WordCount, затем найди слово с максимальным значением
	m := WordCount(text)
	sum := 0
	text = ""
	for k, v := range m {
		if v > sum {
			sum = v
			text = k
		}
	}
	return text
}

func main() {
	text := "go go go python java go python"

	fmt.Println("=== Подсчёт слов ===")
	fmt.Println("Текст:", text)
	fmt.Println("Количество слов:", WordCount(text))
	fmt.Println("Самое частое слово:", MostFrequent(text))

	fmt.Println("\n=== Подсчёт символов ===")
	word := "hello"
	fmt.Println("Слово:", word)
	fmt.Println("Количество символов:")
	for k, v := range CharCount(word) {
		fmt.Printf("%c: %v\n", k, v)
	}

}
