package advertisementCreate

import (
	"net/http"
	"testing"

	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/tests/fixtures"

	"github.com/stretchr/testify/require"
)

func TestCreateAdvertisement(t *testing.T) {
	user := fixtures.Login(t, false)
	fields := fixtures.Fields()
	created := fixtures.Advertisement(t, user.Token, fields)
	id := created.Get("id").String()
	t.Log("Шаг: получение созданного объявления")
	ad := ads.Get(t, user.Token, id, http.StatusOK)
	require.Equal(t, id, ad.Get("id").String(), "Получено другое объявление")
	require.Equal(t, user.User.Id, ad.Get("user_id").String(), "Неверный владелец")
	for _, field := range []string{"title", "description", "price", "quantity"} {
		require.Equal(t, fields[field], ad.Get(field).String(), "Поле %s отличается от отправленного", field)
	}
	t.Log("Шаг: поиск по полному названию")
	list := ads.GetAdvertisementsList(t, user.Token, fields["title"])
	require.Len(t, list.Items, 1, "По уникальному названию должно находиться одно объявление")
	require.Equal(t, id, list.Items[0].ID, "В поиске найдено другое объявление")
	t.Log("Шаг: получение фотографий")
	photos := ads.Photos(t, user.Token, id, http.StatusOK).Array()
	require.Len(t, photos, 1, "Должна быть загружена одна фотография")
	require.NotEmpty(t, photos[0].Get("id").String(), "Нет id фотографии")
	require.NotEmpty(t, photos[0].Get("url").String(), "Нет URL фотографии")
	require.Equal(t, created.Get("photos.0.id").String(), photos[0].Get("id").String(), "Получена другая фотография")
	t.Log("Результат: поля объявления, поиск и фотография проверены")
}
