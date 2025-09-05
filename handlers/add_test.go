package handlers

import (
	"gameCatalogAPI/module"
	"gameCatalogAPI/utils"
	"net/http"
	"testing"
)

func TestAddHandler(t *testing.T) {
	// створення тестових данних
	var test = []module.HandlerTest{
		{Url: "/add?name=HollowKnight&genre=Metroidvania&year=2017&rating=10", ExpectCode: http.StatusOK, Expectlenght: 1},
		{Url: "/add?name=EldenRing&genre=Souls&year=2022&rating=9.5", ExpectCode: http.StatusOK, Expectlenght: 1},
		{Url: "/add?genre=Metroidvania&year=2017&rating=10", ExpectCode: http.StatusBadRequest, Expectlenght: 0},
		{Url: "/add?name=HollowKnight&genre=Metroidvania&year=1940&rating=10", ExpectCode: http.StatusBadRequest, Expectlenght: 0},
		{Url: "/add?name=HollowKnight&genre=Metroidvania&year=2017&rating=abc", ExpectCode: http.StatusBadRequest, Expectlenght: 0},
	}
	// передаччя в функцію
	utils.RunHandlerTest(t, AddHandler, nil, test, http.MethodPost, true)
}
