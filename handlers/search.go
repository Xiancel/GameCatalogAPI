package handlers

import (
	"encoding/json"
	mod "gameCatalogAPI/module"
	"net/http"
	"strconv"
)

// функція/хендлер для пошуку ігор по параметрам
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// запрос параметрів у користувача
	name := r.URL.Query().Get("name")
	genre := r.URL.Query().Get("genre")
	yearStr := r.URL.Query().Get("year")
	ratingStr := r.URL.Query().Get("rating")

	// перевірка на введення одного параметру
	params := make(map[string]string)
	if name != "" {
		params["name"] = name
	}
	if genre != "" {
		params["genre"] = genre
	}
	if yearStr != "" {
		params["year"] = yearStr
	}
	if ratingStr != "" {
		params["rating"] = ratingStr
	}
	if len(params) != 1 {
		http.Error(w, "We apologize, at the moment we only support filtering by one parameter.", http.StatusBadRequest)
		return
	}

	// доставання ключа і зачення для дальнейшего користування
	var field, value string
	for k, v := range params {
		field, value = k, v
	}

	// передача ключа і значення  в функцію фільтрації списку
	slice := filter(value, field)

	// пагинація
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pagined := parginate(slice, page)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagined)
}

// функція пагінації
func parginate(slice []mod.GameCatalog, page int) []mod.GameCatalog {
	limit := 5
	start := (page - 1) * limit
	end := start + limit

	if start > len(slice) {
		start = len(slice)
	}
	if end > len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}

// функція фультрації
func filter(input, param string) []mod.GameCatalog {
	// слайс для зберегання нового списку
	var slice []mod.GameCatalog
	for _, n := range mod.CatalogList {
		// отримання параметру для порівняння
		field := getField(n, param)
		// пороівняння
		if input == field {
			// додавання в слайс
			slice = append(slice, n)
		}
	}
	return slice
}

// функція для взяття параметру
func getField(n mod.GameCatalog, param string) string {
	switch param {
	case "name":
		return n.Name
	case "genre":
		return n.Genre
	case "year":
		return strconv.Itoa(n.Year)
	case "rating":
		return strconv.FormatFloat(n.Rating, 'f', -1, 64)
	default:
		return ""
	}
}
