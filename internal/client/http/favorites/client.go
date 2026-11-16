package favorites

import (
	"net/http"
	"testing"

	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpFavorites sends a request to the user's favorites
func HttpFavorites(t *testing.T, token, method, id, search string) *http.Response {
	endpoint := "/favorites"
	if method == http.MethodDelete {
		endpoint += "/" + id
	}
	request := apiRunner.GetRunner().Auth(token).Create().Method(method).URL(endpoint)
	if method == http.MethodPost {
		request.ContentType("application/json").JSON(map[string]string{"advertisement_id": id})
	}
	if method == http.MethodGet {
		request.Query("search", search).Query("page_size", "50")
	}
	return request.Expect(t).End().Response
}
