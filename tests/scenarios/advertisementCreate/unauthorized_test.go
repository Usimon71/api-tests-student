package advertisementCreate

import (
	ads "api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"net/http"
	"testing"
)

func TestCreateWithoutAuthorization(t *testing.T) {
	base.SetupSuite()
	for _, token := range []string{"", fixtures.InvalidToken} {
		t.Run("token="+token, func(t *testing.T) {
			t.Log("Шаг: создание объявления без авторизации")
			body := ads.Create(t, token, fixtures.Fields(), fixtures.Photos(), http.StatusUnauthorized)
			fixtures.Unauthorized(t, body, token)
		})
	}
}
