package module

import (
	"testing"
)

func TestGameCatalog(t *testing.T) {
	tests := []struct {
		game     GameCatalog
		expected GameCatalog
	}{
		{
			game: GameCatalog{
				ID:     1,
				Name:   "Hollow Knight",
				Genre:  "Metroidvania",
				Year:   2017,
				Rating: 10,
			},
			expected: GameCatalog{
				ID:     1,
				Name:   "Hollow Knight",
				Genre:  "Metroidvania",
				Year:   2017,
				Rating: 10,
			},
		},
		{
			game: GameCatalog{
				ID:     2,
				Name:   "Elden Ring",
				Genre:  "Souls like",
				Year:   2022,
				Rating: 9.6,
			},
			expected: GameCatalog{
				ID:     2,
				Name:   "Elden Ring",
				Genre:  "Souls like",
				Year:   2022,
				Rating: 9.6,
			},
		},
	}

	for i, tt := range tests {
		if tt.game.ID != tt.expected.ID {
			t.Errorf("test %d: expected ID %d, got %d", i, tt.expected.ID, tt.game.ID)
		}
		if tt.game.Name != tt.expected.Name {
			t.Errorf("test %d: expected Name %s, got %s", i, tt.expected.Name, tt.game.Name)
		}
		if tt.game.Genre != tt.expected.Genre {
			t.Errorf("test %d: expected Genre %s, got %s", i, tt.expected.Genre, tt.game.Genre)
		}
		if tt.game.Year != tt.expected.Year {
			t.Errorf("test %d: expected Year %d, got %d", i, tt.expected.Year, tt.game.Year)
		}
		if tt.game.Rating != tt.expected.Rating {
			t.Errorf("test %d: expected Rating %f, got %f", i, tt.expected.Rating, tt.game.Rating)
		}
	}
}

func TestStatsCatalog(t *testing.T) {
	stats := StatsCatalog{
		TotalGame:        8,
		AvgRating:        27.8,
		OldestGame:       "GTA SA",
		NewstGame:        "Hollow Knight Silksong",
		MostPopularGanre: "Souls like",
	}
	expect := StatsCatalog{
		TotalGame:        8,
		AvgRating:        27.8,
		OldestGame:       "GTA SA",
		NewstGame:        "Hollow Knight Silksong",
		MostPopularGanre: "Souls like",
	}
	if stats != expect {
		t.Errorf("Stats do not match expected values")
	}
}
func TestCatalogList(t *testing.T) {
	CatalogList = []GameCatalog{
		{ID: 1, Name: "Game A", Genre: "Action", Year: 2000, Rating: 7.5},
		{ID: 2, Name: "Game B", Genre: "RPG", Year: 2005, Rating: 9.0},
	}

	expected := []GameCatalog{
		{ID: 1, Name: "Game A", Genre: "Action", Year: 2000, Rating: 7.5},
		{ID: 2, Name: "Game B", Genre: "RPG", Year: 2005, Rating: 9.0},
	}

	if len(CatalogList) != 2 {
		t.Errorf("expected 2 games,got %d", len(CatalogList))
	}

	for i, game := range CatalogList {
		if game.Name != expected[i].Name {
			t.Errorf("game %d: expected Name %s, got %s", i, expected[i].Name, game.Name)
		}
		if game.Genre != expected[i].Genre {
			t.Errorf("game %d: expected Genre %s, got %s", i, expected[i].Genre, game.Genre)
		}
		if game.Year != expected[i].Year {
			t.Errorf("game %d: expected Year %d, got %d", i, expected[i].Year, game.Year)
		}
		if game.Rating != expected[i].Rating {
			t.Errorf("game %d: expected Rating %f, got %f", i, expected[i].Rating, game.Rating)
		}
	}

	CatalogList = nil
}
