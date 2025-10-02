package main

import (
	"fmt"
	"math"
)

func main() {
	const dollarRate = 38.5
	const initialAmount = 5000
	const years = 2
	const monthly = 1500
	const bankRate = 0.15

	monts := years * 12
	montlyRate := bankRate / 12

	totalSum := initialAmount + monthly*float64(monts)
	compoundInitial := initialAmount * math.Pow(1+montlyRate, float64(monts))
	EffectOfContributions := montlyRate * (math.Pow(1+montlyRate, float64(monts)) - 1) / montlyRate
	finalAmount := compoundInitial + EffectOfContributions

	interest := finalAmount - totalSum

	fmt.Println("Калькулятор накопичень")
	fmt.Printf("\nПочаткові дані:\n")
	fmt.Printf("- Початкова сума: %.2f грн\n", initialAmount)
	fmt.Printf("- Щомісячні накопичення: %.2f грн\n", monthly)
	fmt.Printf("- Річна ставка: %.1f%%\n", bankRate*100)
	fmt.Printf("- Термін: %d роки (%d місяців)\n", years, monts)
	fmt.Printf("\nРезультати:\n")
	fmt.Printf("- Загальна сума внесків: %.2f грн\n", totalSum)
	fmt.Printf("- Нараховані відсотки: %.2f грн\n", interest)
	fmt.Printf("- Фінальна сума: %.2f грн\n", finalAmount)
	fmt.Printf("\nУ доларах США:\n")
	fmt.Printf("- Загальна сума внесків: $%.2f\n", totalSum/dollarRate)
	fmt.Printf("- Фінальна сума: $%.2f\n", finalAmount/dollarRate)
}
