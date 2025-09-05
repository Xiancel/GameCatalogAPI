package handlers

import (
	"gameCatalogAPI/module"
	"gameCatalogAPI/utils"
	"net/http"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	// створення тестових данних
	module.CatalogList = []module.GameCatalog{
		{ID: 1, Name: "Hollow Knight", Genre: "Metroidvania", Year: 2017, Rating: 10},
		{ID: 2, Name: "Elden Ring", Genre: "Souls like", Year: 2022, Rating: 9.5},
		{ID: 3, Name: "Terraria", Genre: "Sandbox", Year: 2011, Rating: 10},
	}
	// створення тестових запросів
	var test = []module.HandlerTest{
		{Url: "/search?name=Hollow%20Knight", ExpectCode: http.StatusOK, Expectlenght: 1},
		{Url: "/search?genre=Souls%20like", ExpectCode: http.StatusOK, Expectlenght: 1},
		{Url: "/search?year=2011", ExpectCode: http.StatusOK, Expectlenght: 1},
		{Url: "/search?rating=10", ExpectCode: http.StatusOK, Expectlenght: 2},
		{Url: "/search?name=Hollow%20Knight&genre=Metroidvania", ExpectCode: http.StatusBadRequest, Expectlenght: 0},
	}
	// передаччя в функцію
	utils.RunHandlerTest(t, SearchHandler, module.CatalogList, test, http.MethodGet, false)
}
