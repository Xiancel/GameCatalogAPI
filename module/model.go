package module

// структура Каталогу ігор яка містить:
type GameCatalog struct {
	ID     int     `json:"id"`     // айді игри в каталозі
	Name   string  `json:"name"`   // ім'я
	Genre  string  `json:"genre"`  // жанр
	Year   int     `json:"year"`   // рік віпуску
	Rating float64 `json:"rating"` // рейтінг гри
}

type StatsCatalog struct {
	TotalGame        int     `json:"total_game"`
	AvgRating        float64 `json:"avgerage_rating"`
	OldestGame       string  `json:"oldest_game"`
	NewstGame        string  `json:"newst_game"`
	MostPopularGanre string  `json:"most_popular_genre"`
}

// слайс для зберегання всіх ігор в каталозі
var CatalogList []GameCatalog
