package handlers

import (
	"encoding/json"
	mod "gameCatalogAPI/module"
	"net/http"
	"strconv"
)

// функція/хендлер для відображення всого каталогу ігор
func ListHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка на метод
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}

		// паргинация страницы
		limit := 5
		start := (page - 1) * limit
		end := start + limit

		if start > len(mod.CatalogList) {
			start = len(mod.CatalogList)
		}
		if end > len(mod.CatalogList) {
			end = len(mod.CatalogList)
		}
		pagined := mod.CatalogList[start:end]
		// маршелізація всого катологу
		jsonData, err := json.Marshal(pagined)
		// перевірка на помилку
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
