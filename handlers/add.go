package handlers

import (
	"encoding/json"
	mod "gameCatalogAPI/module"
	"net/http"
	"strconv"
)

// функція/хендлер для додавання нової гри в колекцію
func AddHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка на метод
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// запрос у користувача через параметри
	name := r.URL.Query().Get("name")
	genre := r.URL.Query().Get("genre")
	yearStr := r.URL.Query().Get("year")
	ratingStr := r.URL.Query().Get("rating")

	// перевірка на пусте ім'я
	if name == "" {
		http.Error(w, "enter name", http.StatusBadRequest)
		return
	}
	// перевірка на пустий жанр
	if genre == "" {
		http.Error(w, "enter genre", http.StatusBadRequest)
		return
	}
	// перевод з string в int
	year, err := strconv.Atoi(yearStr)
	// валідація року
	if err != nil || year < 1950 || year > 2025 {
		http.Error(w, "invalid year", http.StatusBadRequest)
		return
	}
	// перевід з string в float64
	rating, err := strconv.ParseFloat(ratingStr, 64)
	// валідація рейтингу
	if err != nil || rating < 0 || rating > 10 {
		http.Error(w, "invalid rating", http.StatusBadRequest)
		return
	}

	// додавання до каталогу нову гру
	games := mod.GameCatalog{
		ID:     len(mod.CatalogList) + 1,
		Name:   name,
		Genre:  genre,
		Year:   year,
		Rating: rating,
	}
	mod.CatalogList = append(mod.CatalogList, games)

	// маршелізаця добавленої гри
	jsonData, _ := json.Marshal(games)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
