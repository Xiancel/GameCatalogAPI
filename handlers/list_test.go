package handlers

import (
	"gameCatalogAPI/module"
	"gameCatalogAPI/utils"
	"net/http"
	"testing"
)

func TestListHandler(t *testing.T) {
	// створення тестових данних
	module.CatalogList = []module.GameCatalog{
		{ID: 1, Name: "Hollow Knight", Genre: "Metroidvania", Year: 2017, Rating: 10},
		{ID: 2, Name: "Elden Ring", Genre: "Souls like", Year: 2022, Rating: 9.5},
		{ID: 3, Name: "Terraria", Genre: "Sandbox", Year: 2011, Rating: 10},
		{ID: 4, Name: "Grand Theft Auto: San Andreas", Genre: "Action-adventure, Open world", Year: 2004, Rating: 10},
		{ID: 5, Name: "Left 4 Dead", Genre: "Cooperative First-person shooter, Survival horror", Year: 2008, Rating: 9.8},
		{ID: 6, Name: "Assassin's Creed 4 Black Flag", Genre: "Action-adventure, Stealth, Open world", Year: 2013, Rating: 9.7},
		{ID: 7, Name: "Gang Beasts", Genre: "Party, Fighting", Year: 2017, Rating: 8},
		{ID: 8, Name: "Clone Drone in the Danger Zone", Genre: "Action, Fighting", Year: 2021, Rating: 9},
	}
	// створення тестових запросів
	var test = []module.HandlerTest{
		{Url: "/list?page=1", ExpectCode: http.StatusOK, Expectlenght: 5},
		{Url: "/list?page=2", ExpectCode: http.StatusOK, Expectlenght: 3},
		{Url: "/list", ExpectCode: http.StatusOK, Expectlenght: 5},
		{Url: "/list?page=3", ExpectCode: http.StatusOK, Expectlenght: 0},
	}

	// передаччя в функцію
	utils.RunHandlerTest(t, ListHandler, module.CatalogList, test, http.MethodGet, false)

}
