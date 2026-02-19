package main

import "fmt"

// AddContact добавляет новый контакт в телефонную книгу
func AddContact(book map[string]string, name, phone string) {
	// TODO: реализуй функцию
	// Просто добавь пару имя:телефон в map
	book[name] = phone
}

// GetContact возвращает телефон по имени и флаг существования
func GetContact(book map[string]string, name string) (string, bool) {
	// TODO: реализуй функцию
	// Используй синтаксис value, ok := map[key]
	for k, v := range book {
		if k == name {
			return v, true
		}
	}
	return "", false
}

// UpdateContact обновляет номер существующего контакта
// Возвращает true если контакт найден и обновлён
func UpdateContact(book map[string]string, name, phone string) bool {
	// TODO: реализуй функцию
	// Сначала проверь, существует ли контакт
	for k, _ := range book {
		if k == name {
			book[k] = phone
			return true
		}
	}
	return false
}

// DeleteContact удаляет контакт из книги
// Возвращает true если контакт был удалён
func DeleteContact(book map[string]string, name string) bool {
	// TODO: реализуй функцию
	// Используй delete(map, key)
	for k, _ := range book {
		if k == name {
			delete(book, k)
			return true
		}
	}
	return false
}

// ListContacts выводит все контакты
func ListContacts(book map[string]string) {
	// TODO: реализуй функцию
	// Пройди по map с range и выведи каждый контакт
	for k, v := range book {
		fmt.Println("Name:", k, ", Number:", v)
	}
}

func main() {
	book := make(map[string]string)

	fmt.Println("=== Телефонная книга ===")

	// Добавление
	AddContact(book, "Алексей", "+7-999-123-4567")
	AddContact(book, "Мария", "+7-999-765-4321")
	AddContact(book, "Иван", "+7-999-555-5555")
	fmt.Println("После добавления:")
	ListContacts(book)

	// Получение
	fmt.Println("\nПоиск контакта 'Алексей':")
	if phone, ok := GetContact(book, "Алексей"); ok {
		fmt.Println("Найден:", phone)
	}

	// Обновление
	fmt.Println("\nОбновление номера Алексея:")
	UpdateContact(book, "Алексей", "+7-999-000-0000")
	ListContacts(book)

	// Удаление
	fmt.Println("\nУдаление Марии:")
	DeleteContact(book, "Мария")
	ListContacts(book)
}
