package main

import (
	"fmt"
	"sort"
)

var players []string
var scores map[string]int
var matchesPlayed map[string]int
var wins map[string]int
var losses map[string]int

func registerPlayer(nickname string) bool {
	if playerExists(nickname) {
		return false
	}
	players = append(players, nickname)
	scores[nickname] = 1000
	matchesPlayed[nickname] = 0
	wins[nickname] = 0
	losses[nickname] = 0
	return true
}

func removePlayer(nickname string) bool {
	index := findPlayerIndex(nickname)
	if index == -1 {
		return false
	}
	players = append(players[:index], players[index+1:]...)
	delete(scores, nickname)
	delete(matchesPlayed, nickname)
	delete(wins, nickname)
	delete(losses, nickname)
	return true
}

func findPlayerIndex(nickname string) int {
	for i, name := range players {
		if name == nickname {
			return i
		}
	}
	return -1
}

func playerExists(nickname string) bool {
	_, ok := scores[nickname]
	return ok
}

func displayAllPlayers() {
	if len(players) == 0 {
		fmt.Println("Поки що немає зареєстрованих гравців.")
		return
	}
	fmt.Println("\n--- Список усіх гравців ---")
	for i, p := range players {
		fmt.Printf("%d. %s | Рейтинг: %d | Матчів: %d | Перемог: %d | Поразок: %d | WinRate: %.2f%%\n",
			i+1, p, scores[p], matchesPlayed[p], wins[p], losses[p], calculateWinRate(p))
	}
}

func updateRating(nickname string, won bool, pointsChange int) {
	if !playerExists(nickname) {
		fmt.Println("Гравця не знайдено.")
		return
	}
	matchesPlayed[nickname]++
	if won {
		wins[nickname]++
		scores[nickname] += pointsChange
		fmt.Printf("%s переміг! Рейтинг зріс на %d очок.\n", nickname, pointsChange)
	} else {
		losses[nickname]++
		scores[nickname] -= pointsChange
		if scores[nickname] < 0 {
			scores[nickname] = 0
		}
		fmt.Printf("%s програв. Рейтинг зменшився на %d очок.\n", nickname, pointsChange)
	}
}

func getTopPlayers(count int) []string {
	sorted := sortPlayersByRating()
	if count > len(sorted) {
		count = len(sorted)
	}
	return sorted[:count]
}

func findPlayersByRatingRange(minRating, maxRating int) []string {
	var result []string
	for _, p := range players {
		if scores[p] >= minRating && scores[p] <= maxRating {
			result = append(result, p)
		}
	}
	return result
}

func calculateAverageRating() float64 {
	if len(players) == 0 {
		return 0
	}
	total := 0
	for _, p := range players {
		total += scores[p]
	}
	return float64(total) / float64(len(players))
}

func getBestPlayer() string {
	if len(players) == 0 {
		return ""
	}
	best := players[0]
	for _, p := range players {
		if scores[p] > scores[best] {
			best = p
		}
	}
	return best
}

func getWorstPlayer() string {
	if len(players) == 0 {
		return ""
	}
	worst := players[0]
	for _, p := range players {
		if scores[p] < scores[worst] {
			worst = p
		}
	}
	return worst
}

func calculateWinRate(nickname string) float64 {
	if matchesPlayed[nickname] == 0 {
		return 0
	}
	return float64(wins[nickname]) / float64(matchesPlayed[nickname]) * 100
}

func displayPlayerStats(nickname string) {
	if !playerExists(nickname) {
		fmt.Println("Такого гравця немає.")
		return
	}
	fmt.Printf("\nСтатистика гравця: %s\n", nickname)
	fmt.Printf("Рейтинг: %d | Матчів: %d | Перемог: %d | Поразок: %d | WinRate: %.2f%%\n",
		scores[nickname], matchesPlayed[nickname], wins[nickname], losses[nickname], calculateWinRate(nickname))
}

func displaySystemStats() {
	fmt.Printf("\n=== Загальна статистика ===\n")
	fmt.Printf("Кількість гравців: %d\n", len(players))
	fmt.Printf("Середній рейтинг: %.2f\n", calculateAverageRating())
	if len(players) > 0 {
		fmt.Printf("Найкращий гравець: %s (%d)\n", getBestPlayer(), scores[getBestPlayer()])
		fmt.Printf("Найгірший гравець: %s (%d)\n", getWorstPlayer(), scores[getWorstPlayer()])
	}
}

func displayMenu() {
	fmt.Println("\n=== Система рейтингу гравців ===")
	fmt.Println("1. Зареєструвати гравця")
	fmt.Println("2. Видалити гравця")
	fmt.Println("3. Оновити рейтинг після матчу")
	fmt.Println("4. Знайти гравця")
	fmt.Println("5. Всі гравці")
	fmt.Println("6. Топ гравців")
	fmt.Println("7. Пошук за діапазоном рейтингу")
	fmt.Println("8. Статистика гравця")
	fmt.Println("9. Загальна статистика")
	fmt.Println("10. Вихід")
}

func getStringInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func getIntInput() int {
	var input int
	for {
		_, err := fmt.Scan(&input)
		if err == nil {
			break
		}
		fmt.Println("Некоректне значення, спробуйте ще раз:")
	}
	return input
}

func sortPlayersByRating() []string {
	sorted := make([]string, len(players))
	copy(sorted, players)
	sort.Slice(sorted, func(i, j int) bool {
		return scores[sorted[i]] > scores[sorted[j]]
	})
	return sorted
}

func main() {
	scores = make(map[string]int)
	matchesPlayed = make(map[string]int)
	wins = make(map[string]int)
	losses = make(map[string]int)

	for {
		displayMenu()
		fmt.Print("> ")
		choice := getIntInput()

		switch choice {
		case 1:
			fmt.Print("Введіть нік: ")
			name := getStringInput()
			if registerPlayer(name) {
				fmt.Printf("Гравця %s успішно додано!\n", name)
			} else {
				fmt.Println("Такий гравець вже існує.")
			}
		case 2:
			fmt.Print("Введіть нік: ")
			name := getStringInput()
			if removePlayer(name) {
				fmt.Printf("Гравця %s видалено.\n", name)
			} else {
				fmt.Println("Не знайдено.")
			}
		case 3:
			fmt.Print("Нік: ")
			name := getStringInput()
			fmt.Print("1 - перемога, 0 - поразка: ")
			result := getIntInput()
			fmt.Print("Зміна рейтингу: ")
			points := getIntInput()
			updateRating(name, result == 1, points)
		case 4:
			fmt.Print("Введіть нік: ")
			name := getStringInput()
			displayPlayerStats(name)
		case 5:
			displayAllPlayers()
		case 6:
			fmt.Print("Кількість: ")
			count := getIntInput()
			top := getTopPlayers(count)
			if len(top) == 0 {
				fmt.Println("Немає гравців для відображення.")
			} else {
				fmt.Println("\n--- Топ гравців ---")
				for i, p := range top {
					fmt.Printf("%d. %s - %d\n", i+1, p, scores[p])
				}
			}
		case 7:
			fmt.Print("Мін. рейтинг: ")
			min := getIntInput()
			fmt.Print("Макс. рейтинг: ")
			max := getIntInput()
			res := findPlayersByRatingRange(min, max)
			if len(res) == 0 {
				fmt.Println("Гравців у цьому діапазоні немає.")
			} else {
				fmt.Println("\nГравці в межах рейтингу:")
				for _, p := range res {
					fmt.Printf("%s - %d\n", p, scores[p])
				}
			}
		case 8:
			fmt.Print("Нік: ")
			name := getStringInput()
			displayPlayerStats(name)
		case 9:
			displaySystemStats()
		case 10:
			fmt.Println("До побачення! Гарної гри.")
			return
		default:
			fmt.Println("Невірний вибір. Спробуйте ще раз.")
		}
	}
}
