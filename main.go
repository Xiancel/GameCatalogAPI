package main

import (
	"fmt"
	"gameCatalogAPI/handlers"
	"net/http"
)

// головна функція
func main() {
	// додавання хендлерів
	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/item", handlers.GetItemByIdHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)

	//ініціалізація сервера
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server dead 💀")
		return
	}
}
