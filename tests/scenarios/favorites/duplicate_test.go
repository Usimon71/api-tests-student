package favorites

import (
	"net/http"
	"testing"

	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoriteDuplicate(t *testing.T) {
	owner := fixtures.Login(t, false)
	user := fixtures.Login(t, true)
	require.NotEqual(t, owner.User.Id, user.User.Id, "Нужны разные аккаунты")
	fields := fixtures.Fields()
	ad := fixtures.Advertisement(t, owner.Token, fields)
	id := ad.Get("id").String()
	fav.Add(t, user.Token, id, http.StatusCreated)
	t.Log("Шаг: повторное добавление объявления в избранное")
	body := fav.Add(t, user.Token, id, http.StatusConflict)
	fixtures.APIError(t, body, "already")
	list := fav.List(t, user.Token, fields["title"], http.StatusOK)
	require.EqualValues(t, 1, list.Get("total").Int(), "Повторное добавление создало дубликат")
}
