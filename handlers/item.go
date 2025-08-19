package handlers

import (
	"encoding/json"
	"gameCatalogAPI/module"
	"net/http"
	"strconv"
)

// функція/хендлер для відображення гри по її ID
func GetItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	// перевірка на метод
	if r.Method == "GET" {
		// запрос id у користувача
		idStr := r.URL.Query().Get("id")

		// перевід з string в int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}

		// виклик функції пошуку
		game := SearchById(id)

		// валідація id
		if game.ID == 0 {
			return
		}

		// відображення гри
		jsonData, _ := json.Marshal(game)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

// функція для пошуку гри по її ID
func SearchById(id int) module.GameCatalog {
	for i, n := range module.CatalogList {
		if module.CatalogList[i].ID == id {
			return n
		}
	}
	return module.GameCatalog{}
}
