package favorites

import (
	"net/http"
	"testing"

	ads "api-tests-template/internal/managers/advertisements"
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFavoritesCycle(t *testing.T) {
	owner, buyer := fixtures.LoginOwnerAndBuyer(t)
	fields := fixtures.Fields()
	ad := ads.Create(t, owner.Token, fields, fixtures.Photos(), http.StatusCreated)
	id := ad.Get("id").String()
	require.NotEmpty(t, id, "Нет id созданного объявления")
	deleted := false
	t.Cleanup(func() {
		if !deleted {
			ads.Delete(t, owner.Token, id, http.StatusNoContent)
		}
	})
	t.Log("Шаг: добавление чужого объявления в избранное")
	added := fav.Add(t, buyer.Token, id, http.StatusCreated)
	require.Equal(t, id, added.Get("advertisement.id").String(), "Добавлено другое объявление")
	list := fav.List(t, buyer.Token, fields["title"], http.StatusOK)
	require.EqualValues(t, 1, list.Get("total").Int(), "Объявление не появилось в избранном")
	require.Equal(t, id, list.Get("items.0.advertisement.id").String(), "В избранном другое объявление")
	t.Log("Шаг: удаление из избранного")
	fav.Delete(t, buyer.Token, id, http.StatusNoContent)
	list = fav.List(t, buyer.Token, fields["title"], http.StatusOK)
	requireEmptyFavorites(t, list)
	t.Log("Шаг: повторное добавление и удаление самого объявления")
	fav.Add(t, buyer.Token, id, http.StatusCreated)
	ads.Delete(t, owner.Token, id, http.StatusNoContent)
	deleted = true
	list = fav.List(t, buyer.Token, fields["title"], http.StatusOK)
	requireEmptyFavorites(t, list)
	t.Log("Результат: полный цикл избранного проверен")
}

// requireEmptyFavorites checks that the response contains an empty list and a zero total
func requireEmptyFavorites(t *testing.T, list gjson.Result) {
	t.Helper()
	items := list.Get("items")
	total := list.Get("total")
	require.True(t, items.IsArray(), "Поле items должно быть массивом: %s", list.Raw)
	require.Equal(t, gjson.Number, total.Type, "Поле total должно быть числом: %s", list.Raw)
	require.Empty(t, items.Array(), "Удалённое объявление осталось в избранном")
	require.Zero(t, total.Float(), "Ожидали пустой список избранного, ответ: %s", list.Raw)
}
