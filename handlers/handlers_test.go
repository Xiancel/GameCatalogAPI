package handlers

import (
	"encoding/json"
	"fmt"
	"gameCatalogAPI/module"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	module.CatalogList = nil
	req := httptest.NewRequest(http.MethodPost, "/add?name=HollowKnight&genre=Metroidvania&year=2017&rating=10", nil)
	w := httptest.NewRecorder()

	AddHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if len(module.CatalogList) != 1 {
		t.Errorf("expected length catalog 1, got %d", len(module.CatalogList))
	}

	var game module.GameCatalog
	if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if game.Name != "HollowKnight" || game.Genre != "Metroidvania" || game.Year != 2017 || game.Rating != 10 {
		t.Errorf("unexpected game data: %+v", game)
	}

	fmt.Println(game)
}

func TestGetItemByIdHandler(t *testing.T) {
	module.CatalogList = []module.GameCatalog{
		{ID: 1, Name: "Hollow Knight", Genre: "Metroidvania", Year: 2017, Rating: 10},
	}

	req := httptest.NewRequest(http.MethodGet, "/item?id=1", nil)
	w := httptest.NewRecorder()

	GetItemByIdHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200,got %d", w.Code)
	}

	var game module.GameCatalog
	if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
		t.Errorf("failed to decode responsew: %v", err)
	}

	if game.ID != 1 || game.Name != "Hollow Knight" || game.Genre != "Metroidvania" || game.Year != 2017 || game.Rating != 10 {
		t.Errorf("unexpected game data: %+v", game)
	}

	fmt.Println(game)
	module.CatalogList = nil
}

func TestSearchHandler(t *testing.T) {
	module.CatalogList = []module.GameCatalog{
		{ID: 1, Name: "Hollow Knight", Genre: "Metroidvania", Year: 2017, Rating: 10},
		{ID: 2, Name: "Elden Ring", Genre: "Souls like", Year: 2022, Rating: 9.5},
		{ID: 3, Name: "Terraria", Genre: "Sandbox", Year: 2011, Rating: 10},
	}
	test := []struct {
		url          string
		expectCode   int
		expectlenght int
	}{
		{"/search?name=Hollow%20Knight", http.StatusOK, 1},
		{"/search?genre=Souls%20like", http.StatusOK, 1},
		{"/search?year=2011", http.StatusOK, 1},
		{"/search?rating=10", http.StatusOK, 2},
		{"/search?name=Hollow%20Knight&genre=Metroidvania", http.StatusBadRequest, 0},
	}

	for _, tt := range test {
		req := httptest.NewRequest(http.MethodGet, tt.url, nil)
		w := httptest.NewRecorder()

		SearchHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tt.expectCode {
			t.Errorf("URL %s: expected status %d, got %d", tt.url, tt.expectCode, resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			var result []module.GameCatalog
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Errorf("failed to decode JSON: %v", err)
			}
			if len(result) != tt.expectlenght {
				t.Errorf("URL %s: expected  %d results, got %d", tt.url, tt.expectlenght, len(result))
			}
			fmt.Println("Result: ", result)
		}
	}
	module.CatalogList = nil
}
