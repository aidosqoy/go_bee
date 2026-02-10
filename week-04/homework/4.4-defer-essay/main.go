package main

import "fmt"

/*
TODO: Напиши своё объяснение defer здесь

1. Что такое defer?
	defer - это способ отложения выполнения функции до выхода из функции. Он помогает возвращать функцию в конце.
	Также выполняется при return даже если была паника

2. В каком порядке выполняются несколько defer?
	Несколько defer выполняются в порядке Last In First Out(LIFO),
	то есть первым выполняется послдений defer, и так по очереди снизу вверх

3. Примеры использования defer в реальном коде:
   - Пример 1(Обратный отсчет):
		for i := 1; i <= n; i++ {
		defer fmt.Println(i)
		}

   - Пример 2(Разблокировка мьютекса):
		mu.Lock()
		defer mu.Unlock()

   - Пример 3(Закрытие ресурсов):
		file, _ := os.Open("test.txt")
		defer file.Close()


*/

func main() {
	fmt.Println("Начало функции main")

	defer fmt.Println("Это выполнится последним (defer 1)")
	defer fmt.Println("Это выполнится предпоследним (defer 2)")
	defer fmt.Println("Это выполнится третьим с конца (defer 3)")

	fmt.Println("Конец функции main (но до defer)")
}
