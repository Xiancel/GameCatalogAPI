package module

// структура Каталогу ігор яка містить:
type GameCatalog struct {
	ID     int     `json:"id"`     // айді игри в каталозі
	Name   string  `json:"name"`   // ім'я
	Genre  string  `json:"genre"`  // жанр
	Year   int     `json:"year"`   // рік віпуску
	Rating float64 `json:"rating"` // рейтінг гри
}

// слайс для зберегання всіх ігор в каталозі
var CatalogList []GameCatalog
