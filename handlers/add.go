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
	if r.Method == "POST" {
		// запрос у користувача через параметри
		name := r.URL.Query().Get("name")
		genre := r.URL.Query().Get("genre")
		yearStr := r.URL.Query().Get("year")
		ratingStr := r.URL.Query().Get("rating")

		// перевод з string в int
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			return
		}
		// перевід з string в float64
		rating, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil {
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
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
