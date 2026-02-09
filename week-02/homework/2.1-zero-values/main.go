package main

import "fmt"

func main() {
	// TODO: Объяви переменные всех базовых типов БЕЗ инициализации
	// и выведи их значения

	// Целые числа
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Println("=== Целые числа (signed) ===")
	fmt.Printf("int: %d\n", i)
	fmt.Printf("int8: %d\n", i8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("int64: %d\n", i64)

	// TODO: Добавь остальные типы:
	var uintt uint
	var uintt8 uint8
	var uintt16 uint16
	var uintt32 uint32
	var uintt64 uint64
	var floatt32 float32
	var floatt64 float64
	var boolean bool
	var str string

	fmt.Printf("uint: %d\n", uintt)
	fmt.Printf("uint8: %d\n", uintt8)
	fmt.Printf("uint16: %d\n", uintt16)
	fmt.Printf("uint32: %d\n", uintt32)
	fmt.Printf("uint64: %d\n", uintt64)

	fmt.Println("--FLOAT--")
	fmt.Printf("float32: %v\n", floatt32)
	fmt.Printf("float64: %v\n", floatt64)

	fmt.Println("--BOOLEAN--")
	fmt.Printf("boolean: %v\n", boolean)

	fmt.Println("--STRING--")
	fmt.Printf("string: %v\n", str)

	// TODO: Напиши ответ на вопрос здесь:
	// Почему в Go нет null/nil для базовых типов? Какие проблемы это решает?
	//
	// Ответ: В ГОУ у базовых типов нет null/nil потому, что у них есть нулевое значение. Это делает код проще и безопаснее,
	// предотврощяя баги а также ошибки такие, как NullPointerError.
}
