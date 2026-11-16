package advertisements

import (
	"testing"

	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/response"

	"github.com/tidwall/gjson"
)

// Create creates an advertisement and checks the response status
func Create(t *testing.T, token string, fields map[string]string, photos []string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpCreate(t, token, fields, photos), status)
}

// Get returns an advertisement by ID
func Get(t *testing.T, token, id string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpGet(t, token, id), status)
}

// Delete deletes an advertisement and checks the response status
func Delete(t *testing.T, token, id string, status int) {
	t.Helper()
	response.Read(t, client.HttpDelete(t, token, id), status)
}

// Photos returns advertisement photos
func Photos(t *testing.T, token, id string, status int) gjson.Result {
	t.Helper()
	return response.Read(t, client.HttpPhotos(t, token, id), status)
}
