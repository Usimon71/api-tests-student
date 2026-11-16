package response

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

// Read checks the status and reads the JSON response
func Read(t *testing.T, raw *http.Response, status int) gjson.Result {
	t.Helper()
	defer raw.Body.Close()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать ответ")
	require.Equal(t, status, raw.StatusCode, "Неожиданный HTTP статус, ответ: %s", body)
	if status == http.StatusNoContent {
		return gjson.Result{}
	}
	require.True(t, gjson.ValidBytes(body), "Ответ не JSON: %s", body)
	return gjson.ParseBytes(body)
}
