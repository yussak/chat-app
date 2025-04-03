package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestListMessages_DBError(t *testing.T) {
	original := models.GetMessages
	defer func() { models.GetMessages = original }()

	// モック関数に差し替え
	models.GetMessages = func(channelID string) ([]models.Message, error) {
		return nil, errors.New("DB connection failed")
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages?channel_id=test123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ListMessages(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.Contains(t, rec.Body.String(), "データベースエラー: DB connection failed")
}