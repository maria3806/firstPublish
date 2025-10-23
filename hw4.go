package main

import (
	"fmt"
	"math"
	"time"
)

var (
	standardRate = 1.5
	expressRate = 2.5
	standardPackagingRate = 40.0
	reinforcedPackagingRate = 50.0
	premiumPackagingRate = 75.0
)

func calcPrice() {
	fmt.Println("\nРозрахунок вартості доставки")

	weight := getNumberInput("Введіть вагу посилки (кг): ")

	fmt.Println("\nВиберіть тип доставки:")
	fmt.Println("1. Стандартна")
	fmt.Println("2. Експрес")
	deliveryType := getIntInput("> ")

	distance := getNumberInput("Введіть відстань доставки (км): ")

	fmt.Println("\nВиберіть статус клієнта:")
	fmt.Println("1. Звичайний")
	fmt.Println("2. Постійний")
	clientStatus := getIntInput("> ")

	base := calculateBasePrice(weight, distance)
	add := calculateDeliveryTypePrice(base, deliveryType)
	disc := calculateDiscount(base+add, clientStatus)
	final := calculateFinalPrice(base, add, disc)

	fmt.Printf("\nБазова вартість: %.2f грн\n", base)
	fmt.Printf("Додаткова вартість: %.2f грн\n", add)
	fmt.Printf("Знижка: %.2f грн\n", disc)
	fmt.Printf("Загальна вартість: %.2f грн\n", final)
}

func calcTime() {
	fmt.Println("\nцінка часу доставки")

	distance := getNumberInput("Введіть відстань доставки (км): ")

	fmt.Println("\nВиберіть тип доставки:")
	fmt.Println("1. Стандартна")
	fmt.Println("2. Експрес")
	deliveryType := getIntInput("> ")

	fmt.Println("\nВиберіть погодні умови:")
	fmt.Println("1. Хороші")
	fmt.Println("2. Задовільні")
	fmt.Println("3. Погані")
	weather := getIntInput("> ")

	weekend := getYesNoInput("Сьогодні вихідний день? (так/ні): ")

	base := calculateBaseDeliveryTime(distance, deliveryType)
	delay := addWeatherDelay(base, weather)
	final := calculateFinalDeliveryTime(base, delay, weekend)

	fmt.Printf("\nБазовий час доставки: %.1f днів\n", base)
	fmt.Printf("Затримка через погоду: %.1f днів\n", delay)
	fmt.Printf("Загальний орієнтовний час доставки: %.1f днів\n", final)

	deliveryDate := time.Now().AddDate(0, 0, int(math.Ceil(final)))
	fmt.Printf("Орієнтовна дата прибуття: %s\n", deliveryDate.Format("02 січня 2006"))
}

func calcPackaging() {
	fmt.Println("\n--- Розрахунок пакувальних матеріалів ---")

	length := getNumberInput("Введіть довжину посилки (см): ")
	width := getNumberInput("Введіть ширину посилки (см): ")
	height := getNumberInput("Введіть висоту посилки (см): ")

	fmt.Println("\nВиберіть тип пакувального матеріалу:")
	fmt.Println("1. Стандартний картон")
	fmt.Println("2. Посилений картон з плівкою")
	fmt.Println("3. Преміум пакування")
	materialType := getIntInput("> ")

	mat := calculatePackagingMaterial(length, width, height)
	cost := calculatePackagingCost(mat, materialType)

	fmt.Printf("\nНеобхідна кількість пакувального матеріалу: %.2f м²\n", mat)
	fmt.Printf("Вартість пакувальних матеріалів: %.2f грн\n", cost)
}

func calculateBasePrice(weight, distance float64) float64 {
	return math.Round(weight*distance*0.6*100) / 100
}

func calculateDeliveryTypePrice(base float64, deliveryType int) float64 {
	if deliveryType == 2 {
		return math.Round(base*0.5*100) / 100
	}
	return 0
}

func calculateDiscount(price float64, clientStatus int) float64 {
	if clientStatus == 2 {
		return math.Round(price*0.1*100) / 100
	}
	return 0
}

func calculateFinalPrice(base, add, discount float64) float64 {
	return math.Round((base+add-discount)*100) / 100
}

func calculateBaseDeliveryTime(distance float64, deliveryType int) float64 {
	speed := 60.0
	if deliveryType == 2 {
		speed = 90.0
	}
	return math.Round((distance/speed/8)*100) / 100
}

func addWeatherDelay(base float64, weather int) float64 {
	switch weather {
	case 2:
		return 0.5
	case 3:
		return 1.0
	default:
		return 0
	}
}

func calculateFinalDeliveryTime(base, delay float64, weekend bool) float64 {
	if weekend {
		return base + delay + 1
	}
	return base + delay
}

func calculatePackagingMaterial(length, width, height float64) float64 {
	return math.Round(((2*(length*width+width*height+height*length))/10000)*100) / 100
}

func calculatePackagingCost(amount float64, materialType int) float64 {
	var rate float64
	switch materialType {
	case 1:
		rate = standardPackagingRate
	case 2:
		rate = reinforcedPackagingRate
	case 3:
		rate = premiumPackagingRate
	default:
		rate = standardPackagingRate
	}
	return math.Round(amount*rate*100) / 100
}

func getNumberInput(prompt string) float64 {
	var x float64
	fmt.Print(prompt)
	fmt.Scanln(&x)
	return x
}

func getIntInput(prompt string) int {
	var x int
	fmt.Print(prompt)
	fmt.Scanln(&x)
	return x
}

func getYesNoInput(prompt string) bool {
	var s string
	fmt.Print(prompt)
	fmt.Scanln(&s)
	return s == "так" || s == "Так"
}

func main() {
	for {
		fmt.Println("\n---Калькулятор доставки посилок---")
		fmt.Println("1. Розрахунок вартості доставки")
		fmt.Println("2. Оцінка часу доставки")
		fmt.Println("3. Розрахунок пакувальних матеріалів")
		fmt.Println("4. Вихід")
		fmt.Print("> ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			calcPrice()
		case 2:
			calcTime()
		case 3:
			calcPackaging()
		case 4:
			fmt.Println("До побачення!")
			return
		default:
			fmt.Println("Невірний вибір, спробуйте ще раз.")
		}
	}
}
