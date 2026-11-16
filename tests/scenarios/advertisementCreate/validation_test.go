package advertisementCreate

import (
	"net/http"
	"strings"
	"testing"

	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/tests/fixtures"
)

func TestCreateInvalidAdvertisement(t *testing.T) {
	user := fixtures.Login(t, false)
	cases := []struct {
		name, field, value, message string
		omit, noPhotos              bool
	}{
		{name: "без title", field: "title", omit: true, message: "title is required"},
		{name: "пустой title", field: "title", message: "title is required"},
		{name: "длинный title", field: "title", value: strings.Repeat("a", 51), message: "title must not exceed 50 characters provided 51"},
		{name: "без description", field: "description", omit: true, message: "description is required"},
		{name: "пустой description", field: "description", message: "description is required"},
		{name: "отрицательная цена", field: "price", value: "-1", message: "price must be between 0 and 1000000 inclusive"},
		{name: "нечисловая цена", field: "price", value: "abc", message: "cannot convert string value abc to float"},
		{name: "нечисловое количество", field: "quantity", value: "abc", message: "cannot convert string value to int"},
		{name: "нулевое количество", field: "quantity", value: "0", message: "quantity must be between 1 and 100 inclusive"},
		{name: "отрицательное количество", field: "quantity", value: "-1", message: "quantity must be between 1 and 100 inclusive"},
		{name: "длинное описание", field: "description", value: strings.Repeat("a", 501), message: "description must be between 1 and 500 characters"},
		{name: "без фотографий", noPhotos: true, message: "at least 1 photo is required"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fields := fixtures.Fields()
			if tc.field != "" {
				fields[tc.field] = tc.value
			}
			if tc.omit {
				delete(fields, tc.field)
			}
			photos := fixtures.Photos()
			if tc.noPhotos {
				photos = nil
			}
			t.Log("Шаг: создание объявления с некорректными параметрами")
			body := ads.Create(t, user.Token, fields, photos, http.StatusBadRequest)
			fixtures.Error(t, body, "invalid_request", tc.message)
		})
	}
}
