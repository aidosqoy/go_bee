package main

import (
	"errors"
	"fmt"
)

// Account представляет банковский счёт
type Account struct {
	Owner   string
	Balance float64
}

// NewAccount создаёт новый счёт с начальным балансом
func NewAccount(owner string, initial float64) *Account {
	// TODO: реализуй функцию
	// Создай и верни указатель на Account с заполненными полями
	return &Account{owner, initial}
}

// Deposit пополняет счёт на указанную сумму
func (a *Account) Deposit(amount float64) error {
	// TODO: реализуй метод
	// Проверь, что amount > 0
	if amount > 0 {
		a.Balance += amount
		return nil
	}
	return errors.New("not implemented")
}

// Withdraw снимает деньги со счёта
func (a *Account) Withdraw(amount float64) error {
	// TODO: реализуй метод
	// Проверь, что amount > 0 и достаточно средств
	if amount > 0 && a.Balance >= amount {
		a.Balance -= amount
		return nil
	}
	return errors.New("not implemented")
}

// Transfer переводит деньги на другой счёт
func (a *Account) Transfer(to *Account, amount float64) error {
	// TODO: реализуй метод
	// Используй Withdraw и Deposit
	err := a.Withdraw(amount)
	if err != nil {
		return errors.New("not implemented")
	}
	to.Deposit(amount)
	return nil
}

// Statement возвращает выписку по счёту
func (a Account) Statement() string {
	// TODO: реализуй метод
	// Формат: "Счёт: {Owner}, Баланс: {Balance:.2f}"
	return "Account: " + a.Owner + ", Balance: " + fmt.Sprintf("%f", a.Balance)
}

func main() {
	fmt.Println("=== Банковский счёт ===")

	acc1 := NewAccount("Алексей", 1000)
	acc2 := NewAccount("Мария", 500)

	fmt.Println("Начальное состояние:")
	fmt.Println(acc1.Statement())
	fmt.Println(acc2.Statement())

	// Пополнение
	fmt.Println("\nПополнение счёта Алексея на 500:")
	acc1.Deposit(500)
	fmt.Println(acc1.Statement())

	// Снятие
	fmt.Println("\nСнятие 200 со счёта Алексея:")
	acc1.Withdraw(200)
	fmt.Println(acc1.Statement())

	// Перевод
	fmt.Println("\nПеревод 300 от Алексея к Марии:")
	acc1.Transfer(acc2, 300)
	fmt.Println(acc1.Statement())
	fmt.Println(acc2.Statement())

	// Ошибка — недостаточно средств
	fmt.Println("\nПопытка снять 5000:")
	err := acc1.Withdraw(5000)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
