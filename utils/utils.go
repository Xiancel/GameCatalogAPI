package utils

import (
	"encoding/json"
	"fmt"
	"gameCatalogAPI/module"
	"net/http"
	"net/http/httptest"
	"testing"
)

// допоміжна функція для тестування деяких хендлерів
func RunHandlerTest(t *testing.T, handler http.HandlerFunc, catalog []module.GameCatalog, tests []module.HandlerTest, method string, single bool) {
	// передачя даних в каталог
	module.CatalogList = catalog

	for i, tt := range tests {
		// создання запроса та рекордера
		req := httptest.NewRequest(method, tt.Url, nil)
		w := httptest.NewRecorder()

		// передачя запроса та рекордера в хендлер
		handler(w, req)

		// получаем результат з рекордера
		resp := w.Result()
		defer resp.Body.Close()

		// перевірка кода відповіді
		if resp.StatusCode != tt.ExpectCode {
			t.Errorf("URL %s: expected status %d, got %d", tt.Url, tt.ExpectCode, resp.StatusCode)
		}
		// якщо стату = 200 тоді парсим JSON та перевіряемо данні
		if resp.StatusCode == http.StatusOK {
			// перевірка данніх на одиночні чи масив
			if single {
				var result module.GameCatalog
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Errorf("failed to decode JSON: %v", err)
				}
				// вивід результату
				fmt.Printf("\nResult %d: %v\n", i+1, result)
			} else {
				var result []module.GameCatalog
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Errorf("failed to decode JSON: %v", err)
				}

				// перевірка довжини результату
				if len(result) != tt.Expectlenght {
					t.Errorf("URL %s: expected  %d results, got %d", tt.Url, tt.Expectlenght, len(result))
				}
				// вивід результату
				fmt.Printf("\nResult %d: %v\n", i+1, result)
			}
		}
	}
	// очищяемо каталог
	module.CatalogList = nil
}
