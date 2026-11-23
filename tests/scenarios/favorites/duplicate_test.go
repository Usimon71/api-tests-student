package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoriteDuplicate(t *testing.T) {
	owner, buyer := fixtures.LoginOwnerAndBuyer(t)
	fields := fixtures.Fields()
	ad := fixtures.Advertisement(t, owner.Token, fields)
	id := ad.Get("id").String()
	fav.Add(t, buyer.Token, id, http.StatusCreated)
	t.Log("Шаг: повторное добавление объявления в избранное")
	body := fav.Add(t, buyer.Token, id, http.StatusConflict)
	fixtures.APIError(t, body, "already")
	list := fav.List(t, buyer.Token, fields["title"], http.StatusOK)
	require.EqualValues(t, 1, list.Get("total").Int(), "Повторное добавление создало дубликат")
}
