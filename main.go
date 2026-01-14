package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Subject struct {
	id    int
	name  string
	grade int
	notes string
}

var subjects []Subject
var nextID = 1

func main() {
	subjects = append(subjects, Subject{
		id:    nextID,
		name:  "математика",
		grade: 2,
		notes: "складний предмет, потрібно більше практики з інтегралами",
	})
	nextID++

	subjects = append(subjects, Subject{
		id:    nextID,
		name:  "фізика",
		grade: 8,
		notes: "",
	})
	nextID++

	subjects = append(subjects, Subject{
		id:    nextID,
		name:  "програмування",
		grade: 12,
		notes: "найцікавіший предмет",
	})
	nextID++

	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	http.HandleFunc("/view", view)
	http.HandleFunc("/stats", stats)

	fmt.Println("сервер запущено на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method != "GET" {
		http.Error(w, "помилка: метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "моя картотека предметів\n")
	fmt.Fprintf(w, "предмети (%d):\n", len(subjects))

	for _, s := range subjects {
		fmt.Fprintf(w, "%d. %s - оцінка: %d/12\n", s.id, s.name, s.grade)
	}

	fmt.Fprintf(w, "\nвсього предметів: %d\n", len(subjects))
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method != "POST" {
		http.Error(w, "помилка: метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	name := r.FormValue("name")
	gradeText := r.FormValue("grade")
	notes := r.FormValue("notes")

	if name == "" {
		http.Error(w, "помилка: назва предмету обовʼязкова", http.StatusBadRequest)
		return
	}

	grade, err := strconv.Atoi(gradeText)
	if err != nil || grade < 0 || grade > 12 {
		http.Error(w, "помилка: оцінка повинна бути від 0 до 12", http.StatusBadRequest)
		return
	}

	newSubject := Subject{
		id:    nextID,
		name:  name,
		grade: grade,
		notes: notes,
	}
	nextID++

	subjects = append(subjects, newSubject)

	fmt.Fprintln(w, "предмет додано успішно\n")
	fmt.Fprintf(w, "id: %d\n", newSubject.id)
	fmt.Fprintf(w, "назва: %s\n", newSubject.name)
	fmt.Fprintf(w, "оцінка: %d/12\n", newSubject.grade)
	fmt.Fprintf(w, "нотатки: %s\n", newSubject.notes)
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method != "GET" {
		http.Error(w, "помилка: метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	idText := r.URL.Query().Get("id")
	if idText == "" {
		http.Error(w, "помилка: id параметр обовʼязковий", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "помилка: невірний id", http.StatusBadRequest)
		return
	}

	for _, s := range subjects {
		if s.id == id {
			fmt.Fprintln(w, "перегляд предмету\n")
			fmt.Fprintf(w, "id: %d\n", s.id)
			fmt.Fprintf(w, "назва: %s\n", s.name)
			fmt.Fprintf(w, "оцінка: %d/12\n", s.grade)
			fmt.Fprintf(w, "нотатки: %s\n", s.notes)
			return
		}
	}

	http.Error(w, "помилка: предмет не знайдено", http.StatusNotFound)
}

func stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method != "GET" {
		http.Error(w, "помилка: метод не дозволений", http.StatusMethodNotAllowed)
		return
	}

	if len(subjects) == 0 {
		fmt.Fprintln(w, "дані відсутні")
		return
	}

	sum := 0
	best := subjects[0]
	worst := subjects[0]

	for _, s := range subjects {
		sum += s.grade
		if s.grade > best.grade {
			best = s
		}
		if s.grade < worst.grade {
			worst = s
		}
	}

	avg := float64(sum) / float64(len(subjects))

	fmt.Fprintln(w, "статистика навчання\n")
	fmt.Fprintf(w, "всього предметів: %d\n", len(subjects))
	fmt.Fprintf(w, "середній бал: %.1f/12\n", avg)
	fmt.Fprintf(w, "найкраща оцінка: %d/12 (%s)\n", best.grade, best.name)
	fmt.Fprintf(w, "найгірша оцінка: %d/12 (%s)\n", worst.grade, worst.name)
}
