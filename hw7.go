package main

import (
	"fmt"
)

type PaymentMethod interface {
	Validate(amount float64) error
	CalculateFee(amount float64) float64
	Process(amount float64) error
	Name() string
}

func EnterAmount() float64 {
	var amount float64
	fmt.Print("Введіть суму платежу: ")
	fmt.Scanln(&amount)
	return amount
}

type CreditCard struct{}

func (c CreditCard) Validate(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("сума має бути більшою за 0")
	}
	return nil
}

func (c CreditCard) CalculateFee(amount float64) float64 {
	return amount * 0.025
}

func (c CreditCard) Process(amount float64) error {
	return nil
}

func (c CreditCard) Name() string {
	return "Кредитна картка"
}

type PayPal struct{}

func (p PayPal) Validate(amount float64) error {
	if amount < 10 {
		return fmt.Errorf("мінімальна сума для PayPal — 10 грн")
	}
	return nil
}

func (p PayPal) CalculateFee(amount float64) float64 {
	return amount * 0.035
}

func (p PayPal) Process(amount float64) error {
	return nil
}

func (p PayPal) Name() string {
	return "PayPal"
}

type Cash struct{}

func (c Cash) Validate(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("некоректна сума")
	}
	return nil
}

func (c Cash) CalculateFee(amount float64) float64 {
	return 0
}

func (c Cash) Process(amount float64) error {
	return nil
}

func (c Cash) Name() string {
	return "Готівка"
}

type BankTransfer struct{}

func (b BankTransfer) Validate(amount float64) error {
	if amount < 100 {
		return fmt.Errorf("мінімальна сума для банківського переказу — 100 грн")
	}
	return nil
}

func (b BankTransfer) CalculateFee(amount float64) float64 {
	return 15.0
}

func (b BankTransfer) Process(amount float64) error {
	return nil
}

func (b BankTransfer) Name() string {
	return "Банківський переказ"
}

func main() {

	methods := []PaymentMethod{
		CreditCard{},
		PayPal{},
		Cash{},
		BankTransfer{},
	}

	fmt.Println("=== Система платежів ===\n")

	for {
		fmt.Println("Доступні методи оплати:")
		fmt.Println("1. Кредитна картка")
		fmt.Println("2. PayPal")
		fmt.Println("3. Готівка")
		fmt.Println("4. Банківський переказ")
		fmt.Println("5. Вихід\n")

		var choice int
		fmt.Print("Оберіть метод (1-5): ")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Вихід...")
			return
		}

		if choice < 1 || choice > len(methods) {
			fmt.Println("Помилка: некоректний вибір\n")
			continue
		}

		method := methods[choice-1]
		amount := EnterAmount()

		fmt.Printf("\nОбробляємо платіж на суму %.2f грн...\n\n", amount)

		if err := method.Validate(amount); err != nil {
			fmt.Println("Помилка:", err, "\n")
			continue
		}

		fee := method.CalculateFee(amount)
		total := amount + fee

		if err := method.Process(amount); err != nil {
			fmt.Println("Платіж не завершено\n")
			continue
		}

		fmt.Println("Метод:", method.Name())
		fmt.Println("Платіж успішно оброблено")
		fmt.Printf("Сума: %.2f грн\n", amount)
		fmt.Printf("Комісія: %.2f грн\n", fee)
		fmt.Printf("До списання: %.2f грн\n\n", total)
	}
}
