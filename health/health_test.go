package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/health", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, check(c)) {
		assert.Equal(t, rec.Code, http.StatusOK)
		assert.Equal(t, rec.Body.String(), `{"success":true}`)
	}
}
