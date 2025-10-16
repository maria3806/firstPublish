package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Аналізатор погоди для туристів")
		fmt.Println("Виберіть опцію:")
		fmt.Println("1. Аналіз погодних умов")
		fmt.Println("2. Рекомендації для активностей")
		fmt.Println("3. Перевірка безпеки подорожі")
		fmt.Println("0. Вихід")
		fmt.Print("Ваш вибір: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			analyzeWeather()
		case "2":
			activityRecommendations()
		case "3":
			safetyCheck()
		case "0":
			fmt.Println("До побачення!")
			return
		default:
			fmt.Println("Невірний вибір. Спробуйте ще раз.")
		}
	}
}

func getWeatherInput() (int, int, int, string, string, string, bool) {
	var tempInput, humidityInput, windInput, rainInput, timeOfDay, season string

	fmt.Print("Введіть температуру (°C): ")
	fmt.Scanln(&tempInput)
	fmt.Print("Введіть вологість (%): ")
	fmt.Scanln(&humidityInput)
	fmt.Print("Введіть швидкість вітру (м/с): ")
	fmt.Scanln(&windInput)
	fmt.Print("Введіть наявність опадів (так/ні): ")
	fmt.Scanln(&rainInput)
	fmt.Print("Введіть час доби (ранок/день/вечір/ніч): ")
	fmt.Scanln(&timeOfDay)
	fmt.Print("Введіть сезон (весна/літо/осінь/зима): ")
	fmt.Scanln(&season)

	temp, err1 := strconv.Atoi(tempInput)
	humidity, err2 := strconv.Atoi(humidityInput)
	wind, err3 := strconv.Atoi(windInput)
	rain := strings.ToLower(rainInput)
	timeOfDay = strings.ToLower(timeOfDay)
	season = strings.ToLower(season)

	if err1 != nil || err2 != nil || err3 != nil || (rain != "так" && rain != "ні") {
		fmt.Println("Помилка введення. Перевірте значення.")
		return 0, 0, 0, "", "", "", false
	}

	return temp, humidity, wind, rain, timeOfDay, season, true
}

func analyzeWeather() {
	temp, humidity, wind, rain, _, _, ok := getWeatherInput()
	if !ok {
		return
	}

	fmt.Println("\nРезультат:")

	if temp >= 18 && temp <= 27 && humidity >= 30 && humidity <= 70 && wind <= 5 && rain == "ні" {
		fmt.Println("Погодні умови: комфортні")
		fmt.Println("Ризик сонячного удару: низький")
		fmt.Println("Рекомендована активність: пляжний відпочинок")
	} else if temp >= 30 && humidity >= 70 && wind <= 2 && rain == "ні" {
		fmt.Println("Попередження:")
		fmt.Println("Висока ймовірність теплового удару")
		fmt.Println("Уникайте активності з 11:00 до 16:00")
		fmt.Println("Одягайте легкий світлий одяг")
		fmt.Println("Постійно пийте воду")
	} else if rain == "так" && humidity >= 80 {
		fmt.Println("Рекомендації:")
		fmt.Println("Відвідайте музеї")
		fmt.Println("Заплануйте культурну програму")
		fmt.Println("Необхідний водонепроникний одяг")
	} else {
		fmt.Println("Погодні умови: змінні")
		fmt.Println("Уточніть прогноз перед плануванням активностей")
	}
}

func activityRecommendations() {
	temp, humidity, wind, rain, _, season, ok := getWeatherInput()
	if !ok {
		return
	}

	fmt.Println("\nРекомендації:")

	if rain == "так" || wind > 8 {
		fmt.Println("Необхідний водонепроникний одяг")
		fmt.Println("Відвідайте музеї")
		fmt.Println("Культурна програма в приміщенні")
		return
	}

	if temp >= 20 && temp <= 30 && humidity <= 70 && wind <= 5 && rain == "ні" {
		fmt.Println("Пішохідні екскурсії")
		fmt.Println("Прогулянки на природі")
		fmt.Println("Активний відпочинок")
	} else if temp >= 15 && temp < 20 {
		fmt.Println("Прогулянки в парках")
		fmt.Println("Культурна програма")
	} else if temp > 30 {
		fmt.Println("Занадто спекотно для активного відпочинку")
		fmt.Println("Рекомендовано: пляж, басейн, перебування в тіні")
	} else if temp < 10 {
		fmt.Println("Холодна погода")
		fmt.Println("Рекомендовано: музеї, виставки, закриті простори")
	}

	switch season {
	case "літо":
		fmt.Println("Сезон: літо — гарний час для відпочинку на природі")
	case "зима":
		fmt.Println("Сезон: зима — можлива слизька дорога, обережність")
	case "весна":
		fmt.Println("Сезон: весна — змінна погода, слідкуйте за прогнозом")
	case "осінь":
		fmt.Println("Сезон: осінь — можливі дощі, вдягайтесь тепліше")
	default:
		fmt.Println("Невідомий сезон")
	}
}

func safetyCheck() {
	temp, humidity, wind, rain, timeOfDay, season, ok := getWeatherInput()
	if !ok {
		return
	}

	fmt.Println("\nПопередження:")

	if temp > 35 {
		fmt.Println("Надзвичайна спека — високий ризик теплового удару")
	}

	if temp < 0 {
		fmt.Println("Морозна погода — небезпека обмороження")
	}

	if rain == "так" && wind >= 10 {
		fmt.Println("Штормове попередження — залишайтесь у безпечному місці")
	}

	if humidity >= 85 {
		fmt.Println("Висока вологість — дискомфорт під час активностей")
	}

	if wind >= 15 {
		fmt.Println("Сильний вітер — ризик падіння гілок, уникайте лісу")
	}

	if timeOfDay == "ніч" {
		fmt.Println("Обмежена видимість уночі — візьміть ліхтарик")
	}

	switch season {
	case "зима":
		fmt.Println("Одягайтесь тепло — рекомендуються шапка, рукавички")
	case "весна":
		fmt.Println("Перевіряйте прогноз — можливі раптові дощі")
	case "осінь":
		fmt.Println("Теплий одяг — різке похолодання ввечері")
	default:
		fmt.Println("Сезон не розпізнано")
	}
}
