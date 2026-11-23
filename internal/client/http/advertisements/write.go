package advertisements

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpCreate submits an advertisement with form fields and photos
func HttpCreate(t *testing.T, token string, fields map[string]string, photos []string) *http.Response {
	request := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)
	for name, value := range fields {
		request.MultipartFormData(name, value)
	}
	if len(photos) > 0 {
		request.MultipartFile("photos", photos...)
	}
	return request.Expect(t).End().Response
}

// HttpGet fetches an advertisement by ID
func HttpGet(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementPath).Query("id", id).Expect(t).End().Response
}

// HttpDelete deletes an advertisement by ID
func HttpDelete(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).Query("id", id).Expect(t).End().Response
}

// HttpPhotos fetches advertisement photos
func HttpPhotos(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementsPath + "/" + id + "/photos").Expect(t).End().Response
}
