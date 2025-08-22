package handlers

import (
	"encoding/json"
	"gameCatalogAPI/module"
	"net/http"
)

// функція/хендлер для відображення статистики котологу
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка на метод
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// створення статистики ї відображення
	games := module.StatsCatalog{
		TotalGame:        len(module.CatalogList),
		AvgRating:        avgRating(),
		OldestGame:       oldest(),
		NewstGame:        newest(),
		MostPopularGanre: popularGanre(),
	}

	jsonData, _ := json.Marshal(games)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// функція для підрахунку середьного рейтингу всіх ігор
func avgRating() float64 {
	var rating float64
	for _, n := range module.CatalogList {
		rating += n.Rating
	}

	avg := rating / float64(len(module.CatalogList))

	return avg
}

// функція для визначення самої старої гри в каталозі
func oldest() string {
	oldest := module.CatalogList[0].Year
	oldestName := module.CatalogList[0].Name

	for _, n := range module.CatalogList {
		if n.Year < oldest {
			oldestName = n.Name
			oldest = n.Year
		}
	}
	return oldestName
}

// функція для визначення новіїшої гри в каталозі
func newest() string {
	newest := module.CatalogList[0].Year
	newestName := module.CatalogList[0].Name

	for _, n := range module.CatalogList {
		if n.Year > newest {
			newestName = n.Name
			newest = n.Year
		}
	}
	return newestName
}

// функція для визначення популярного жанру в каталозі
func popularGanre() string {
	popularGanre := make(map[string]int)

	for _, n := range module.CatalogList {
		popularGanre[n.Genre]++
	}
	maxCount := 0
	mostPopular := ""
	for genre, count := range popularGanre {
		if count > maxCount {
			maxCount = count
			mostPopular = genre
		}
	}
	return mostPopular
}
