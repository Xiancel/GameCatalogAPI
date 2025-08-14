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
	http.HandleFunc("/aff", handlers.AddHandler)

	//ініціалізація сервера
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server dead 💀")
		return
	}
}
