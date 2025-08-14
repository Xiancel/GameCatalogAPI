package handlers

import (
	"encoding/json"
	mod "gameCatalogAPI/module"
	"net/http"
)

// функція/хендлер для відображення всого каталогу ігор
func ListHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка на метод
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		// маршелізація всого катологу
		jsonData, err := json.Marshal(mod.CatalogList)
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
