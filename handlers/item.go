package handlers

import (
	"encoding/json"
	"gameCatalogAPI/module"
	"net/http"
	"strconv"
)

func GetItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idStr := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}

		game := SearchById(id)

		if game.ID == 0 {
			return
		}

		jsonData, _ := json.Marshal(game)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func SearchById(id int) module.GameCatalog {
	for i, n := range module.CatalogList {
		if module.CatalogList[i].ID == id {
			return n
		}
	}
	return module.GameCatalog{}
}
