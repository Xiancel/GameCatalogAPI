package handlers

import (
	"encoding/json"
	"gameCatalogAPI/module"
	"net/http"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		games := module.StatsCatalog{
			TotalGame:        len(module.CatalogList),
			AvgRating:        avgRating(),
			OldestGame:       oldest(),
			NewstGame:        newest(),
			MostPopularGanre: popularGanre(),
		}

		jsonData, _ := json.Marshal(games)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func avgRating() float64 {
	var rating float64
	for _, n := range module.CatalogList {
		rating += n.Rating
	}

	avg := rating / float64(len(module.CatalogList))

	return avg
}

func oldest() string {
	oldest := module.CatalogList[0].Year
	oldestName := module.CatalogList[0].Name

	for _, n := range module.CatalogList {
		if n.Year < oldest {
			oldestName = n.Name
			oldest = n.Year
		}
	}
	return oldestName
}

func newest() string {
	newest := module.CatalogList[0].Year
	newestName := module.CatalogList[0].Name

	for _, n := range module.CatalogList {
		if n.Year > newest {
			newestName = n.Name
			newest = n.Year
		}
	}
	return newestName
}

func popularGanre() string {
	popularGanre := make(map[string]int)

	for _, n := range module.CatalogList {
		popularGanre[n.Genre]++
	}
	maxCount := 0
	mostPopular := ""
	for genre, count := range popularGanre {
		if count > maxCount {
			maxCount = count
			mostPopular = genre
		}
	}
	return mostPopular
}
