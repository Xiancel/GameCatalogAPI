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
