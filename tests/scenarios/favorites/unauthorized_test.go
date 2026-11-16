package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"

	"github.com/tidwall/gjson"
)

func TestFavoritesWithoutAuthorization(t *testing.T) {
	base.SetupSuite()
	for _, token := range []string{"", fixtures.InvalidToken} {
		for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodDelete} {
			t.Run(method+"/token="+token, func(t *testing.T) {
				t.Log("Шаг: обращение к избранному без авторизации")
				var body gjson.Result
				switch method {
				case http.MethodGet:
					body = fav.List(t, token, "", http.StatusUnauthorized)
				case http.MethodPost:
					body = fav.Add(t, token, fixtures.MissingID, http.StatusUnauthorized)
				case http.MethodDelete:
					body = fav.Delete(t, token, fixtures.MissingID, http.StatusUnauthorized)
				}
				fixtures.Unauthorized(t, body, token)
			})
		}
	}
}
