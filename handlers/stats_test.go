package handlers

import (
	"encoding/json"
	"fmt"
	"gameCatalogAPI/module"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatsHandler(t *testing.T) {
	// створення тестових данних
	module.CatalogList = []module.GameCatalog{
		{ID: 1, Name: "Hollow Knight", Genre: "Metroidvania", Year: 2017, Rating: 10},
		{ID: 2, Name: "Elden Ring", Genre: "Souls like", Year: 2022, Rating: 9.6},
		{ID: 3, Name: "Terraria", Genre: "Sandbox", Year: 2011, Rating: 9.8},
		{ID: 4, Name: "Grand Theft Auto: San Andreas", Genre: "Action-adventure, Open world", Year: 2004, Rating: 9.3},
		{ID: 5, Name: "Left 4 Dead", Genre: "Cooperative First-person shooter, Survival horror", Year: 2008, Rating: 8.6},
		{ID: 6, Name: "Assassin's Creed 4 Black Flag", Genre: "Action-adventure, Open world", Year: 2013, Rating: 9.7},
		{ID: 7, Name: "Gang Beasts", Genre: "Party, Fighting", Year: 2017, Rating: 8},
		{ID: 8, Name: "Clone Drone in the Danger Zone", Genre: "Action, Fighting", Year: 2021, Rating: 9},
	}
	// создання запроса та рекордера
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	// передачя запроса та рекордера в хендлер
	StatsHandler(w, req)

	// получаем результат з рекордера
	resp := w.Result()
	defer resp.Body.Close()

	// перевірка кода відповіді
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	// парсим JSON та перевіряемо данні
	var stats module.StatsCatalog
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}
	// перевірка данних
	if stats.TotalGame != 8 || stats.AvgRating != 9.25 || stats.OldestGame != "Grand Theft Auto: San Andreas" || stats.NewstGame != "Elden Ring" || stats.MostPopularGanre != "Action-adventure, Open world" {
		t.Errorf("unexpected stats data: %+v", stats)
	}

	// Вивід результату
	fmt.Println(stats)

	//очищення каталогу
	module.CatalogList = nil
}
