package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoritesLimit(t *testing.T) {
	owner := fixtures.Login(t, false)
	user := fixtures.Login(t, true)
	require.NotEqual(t, owner.User.Id, user.User.Id, "Нужны разные аккаунты")
	list := fav.List(t, user.Token, "", http.StatusOK)
	count := int(list.Get("total").Int())
	require.LessOrEqual(t, count, fixtures.FavoritesLimit, "Превышен лимит избранного до теста")
	t.Log("Подготовка: заполнение избранного до лимита")
	for i := count; i < fixtures.FavoritesLimit; i++ {
		ad := fixtures.Advertisement(t, owner.Token, fixtures.Fields())
		fav.Add(t, user.Token, ad.Get("id").String(), http.StatusCreated)
	}
	fields := fixtures.Fields()
	extra := fixtures.Advertisement(t, owner.Token, fields)
	t.Log("Шаг: добавление объявления сверх лимита")
	body := fav.Add(t, user.Token, extra.Get("id").String(), http.StatusBadRequest)
	fixtures.APIError(t, body, "limit")
	require.Empty(t, fav.List(t, user.Token, fields["title"], http.StatusOK).Get("items").Array(), "Объявление сверх лимита добавилось")
}
