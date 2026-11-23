package fixtures

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	ads "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

const (
	InvalidToken   = "incorrect_token"
	MissingID      = "00000000-0000-4000-8000-000000000000"
	Description    = "Тестовое объявление для лабораторной работы"
	Price          = "123.45"
	Quantity       = "3"
	FavoritesLimit = 50
)

// Login authenticates the primary or secondary test account
func Login(t *testing.T, second bool) models.LoginOkResponse {
	t.Helper()
	base.SetupSuite()
	prefix := "TEST_"
	if second {
		prefix = "TEST_SECOND_"
	}
	t.Log("Подготовка: авторизация тестового пользователя")
	return auth.Login(t, os.Getenv(prefix+"LOGIN"), os.Getenv(prefix+"PASSWORD"))
}

// LoginOwnerAndBuyer authenticates two distinct accounts for favorites tests
func LoginOwnerAndBuyer(t *testing.T) (models.LoginOkResponse, models.LoginOkResponse) {
	t.Helper()
	owner := Login(t, false)
	buyer := Login(t, true)
	require.NotEqual(t, owner.User.Id, buyer.User.Id, "Для избранного нужны разные аккаунты")
	return owner, buyer
}

// Fields returns all advertisement fields with a unique title
func Fields() map[string]string {
	return map[string]string{
		"title":       fmt.Sprintf("Лабораторная %d", time.Now().UnixNano()),
		"description": Description, "price": Price, "quantity": Quantity,
	}
}

// Photos returns the path to the test PNG image
func Photos() []string {
	_, file, _, _ := runtime.Caller(0)
	return []string{filepath.Join(filepath.Dir(file), "../testdata/photo.png")}
}

// Advertisement creates an advertisement and registers its cleanup
func Advertisement(t *testing.T, token string, fields map[string]string) gjson.Result {
	t.Helper()
	t.Log("Подготовка: создание объявления")
	ad := ads.Create(t, token, fields, Photos(), http.StatusCreated)
	id := ad.Get("id").String()
	require.NotEmpty(t, id, "Созданное объявление должно иметь id")
	t.Cleanup(func() {
		t.Log("Очистка: удаление тестового объявления")
		ads.Delete(t, token, id, http.StatusNoContent)
	})
	return ad
}

// Error checks the error code and server message
func Error(t *testing.T, body gjson.Result, code, message string) {
	t.Helper()
	require.Equal(t, code, body.Get("error").String(), "Неверный код ошибки: %s", body.Raw)
	require.Equal(t, message, body.Get("message").String(), "Неверное сообщение: %s", body.Raw)
	t.Log("Проверка: статус и сообщение об ошибке соответствуют ожиданиям")
}

// Unauthorized checks the message for a missing or invalid token
func Unauthorized(t *testing.T, body gjson.Result, token string) {
	t.Helper()
	message := "Authorization header required"
	if token != "" {
		message = "Invalid or expired token"
	}
	Error(t, body, "unauthorized", message)
}

// APIError checks the error code presence and message content
func APIError(t *testing.T, body gjson.Result, messagePart string) {
	t.Helper()
	require.NotEmpty(t, body.Get("error").String(), "Нет кода ошибки: %s", body.Raw)
	message := body.Get("message").String()
	require.NotEmpty(t, message, "Нет сообщения об ошибке: %s", body.Raw)
	require.Contains(t, strings.ToLower(message), messagePart, "Сообщение не объясняет ожидаемую ошибку: %s", body.Raw)
	t.Logf("Результат: получена ожидаемая ошибка: %s", body.Raw)
}
