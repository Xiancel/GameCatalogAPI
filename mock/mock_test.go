package mock

import (
	"errors"
	"gameCatalogAPI/module"
	"testing"
)

type CatalogRepository interface {
	GetAll() []module.GameCatalog
	GetByID(id int) (*module.GameCatalog, error)
}

type MockCatalog struct {
	Data []module.GameCatalog
}

func (m *MockCatalog) GetAll() []module.GameCatalog {
	return m.Data
}

func (m *MockCatalog) GetByID(id int) (*module.GameCatalog, error) {
	for _, g := range m.Data {
		if g.ID == id {
			return &g, nil
		}
	}
	return nil, errors.New("not found")
}

func TestListHandlerWithMock(t *testing.T) {
	mockRepo := &MockCatalog{
		Data: []module.GameCatalog{
			{ID: 1, Name: "Game A", Genre: "Action", Year: 2000, Rating: 7.5},
			{ID: 2, Name: "Game B", Genre: "RPG", Year: 2005, Rating: 9.0},
		},
	}

	game, err := mockRepo.GetByID(1)
	if err != nil {
		t.Errorf("expected to find game,error: %v", err)
	}
	if game.Name != "Game A" {
		t.Errorf("expected game Game A,got %s", game.Name)
	}

	_, err = mockRepo.GetByID(999)
	if err == nil {
		t.Errorf("expected error for non-existent game,got nil")
	}

	all := mockRepo.GetAll()
	if len(all) != 2 {
		t.Errorf("expecred 2 games, got %d", len(all))
	}
}
