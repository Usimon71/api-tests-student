package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesMissingAdvertisement(t *testing.T) {
	user := fixtures.Login(t, false)
	for _, method := range []string{http.MethodPost, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			t.Log("Шаг: операция с несуществующим объявлением")
			body := fav.Delete
			if method == http.MethodPost {
				body = fav.Add
			}
			result := body(t, user.Token, fixtures.MissingID, http.StatusNotFound)
			fixtures.APIError(t, result, "not found")
		})
	}
}
