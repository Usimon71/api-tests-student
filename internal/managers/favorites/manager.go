package favorites

import (
	"net/http"
	"testing"

	client "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/response"

	"github.com/tidwall/gjson"
)

// Add adds an advertisement to favorites
func Add(t *testing.T, token, id string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpFavorites(t, token, http.MethodPost, id, ""), status)
}

// Delete removes an advertisement from favorites
func Delete(t *testing.T, token, id string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpFavorites(t, token, http.MethodDelete, id, ""), status)
}

// List returns favorites filtered by title
func List(t *testing.T, token, search string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpFavorites(t, token, http.MethodGet, "", search), status)
}
