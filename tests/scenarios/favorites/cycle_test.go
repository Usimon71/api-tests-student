package favorites

import (
	"net/http"
	"testing"

	ads "api-tests-template/internal/managers/advertisements"
	fav "api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestFavoritesCycle(t *testing.T) {
	owner := fixtures.Login(t, false)
	user := fixtures.Login(t, true)
	require.NotEqual(t, owner.User.Id, user.User.Id, "Для избранного нужны разные аккаунты")
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
	added := fav.Add(t, user.Token, id, http.StatusCreated)
	require.Equal(t, id, added.Get("advertisement.id").String(), "Добавлено другое объявление")
	list := fav.List(t, user.Token, fields["title"], http.StatusOK)
	require.EqualValues(t, 1, list.Get("total").Int(), "Объявление не появилось в избранном")
	require.Equal(t, id, list.Get("items.0.advertisement.id").String(), "В избранном другое объявление")
	t.Log("Шаг: удаление из избранного")
	fav.Delete(t, user.Token, id, http.StatusNoContent)
	list = fav.List(t, user.Token, fields["title"], http.StatusOK)
	require.Empty(t, list.Get("items").Array(), "Удалённое объявление осталось в избранном")
	t.Log("Шаг: повторное добавление и удаление самого объявления")
	fav.Add(t, user.Token, id, http.StatusCreated)
	ads.Delete(t, owner.Token, id, http.StatusNoContent)
	deleted = true
	list = fav.List(t, user.Token, fields["title"], http.StatusOK)
	require.Empty(t, list.Get("items").Array(), "Удалённое объявление осталось в избранном")
	require.Zero(t, list.Get("total").Int(), "Неверное число избранных после удаления объявления")
	t.Log("Результат: полный цикл избранного проверен")
}
