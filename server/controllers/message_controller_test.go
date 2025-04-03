package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListMessages_EmptyChannelID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ListMessages(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "ChannelIDが必要です", rec.Body.String())
}
