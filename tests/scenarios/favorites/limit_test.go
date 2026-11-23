package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoritesLimit(t *testing.T) {
	owner, buyer := fixtures.LoginOwnerAndBuyer(t)
	list := fav.List(t, buyer.Token, "", http.StatusOK)
	count := int(list.Get("total").Int())
	require.LessOrEqual(t, count, fixtures.FavoritesLimit, "Превышен лимит избранного до теста")
	t.Log("Подготовка: заполнение избранного до лимита")
	for i := count; i < fixtures.FavoritesLimit; i++ {
		ad := fixtures.Advertisement(t, owner.Token, fixtures.Fields())
		fav.Add(t, buyer.Token, ad.Get("id").String(), http.StatusCreated)
	}
	fields := fixtures.Fields()
	extra := fixtures.Advertisement(t, owner.Token, fields)
	t.Log("Шаг: добавление объявления сверх лимита")
	body := fav.Add(t, buyer.Token, extra.Get("id").String(), http.StatusBadRequest)
	fixtures.APIError(t, body, "limit")
	require.Empty(t, fav.List(t, buyer.Token, fields["title"], http.StatusOK).Get("items").Array(), "Объявление сверх лимита добавилось")
}
