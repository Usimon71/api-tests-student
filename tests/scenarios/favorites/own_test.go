package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoriteOwnAdvertisement(t *testing.T) {
	user := fixtures.Login(t, false)
	fields := fixtures.Fields()
	ad := fixtures.Advertisement(t, user.Token, fields)
	t.Log("Шаг: добавление своего объявления в избранное")
	body := fav.Add(t, user.Token, ad.Get("id").String(), http.StatusForbidden)
	fixtures.APIError(t, body, "own")
	require.Empty(t, fav.List(t, user.Token, fields["title"], http.StatusOK).Get("items").Array(), "Своё объявление попало в избранное")
}
