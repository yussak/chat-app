package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"server/models"
	"strings"
	"testing"
	"time"

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

func TestListMessages_success(t *testing.T) {
	original := models.GetMessages
	defer func() { models.GetMessages = original }()

	mockMessages := []models.Message{
		{
			ID:        1,
			Content:   "test message",
			CreatedAt: time.Date(2025, 4, 3, 21, 58, 47, 107456000, time.FixedZone("JST", 9*60*60)),
			ChannelID: 1,
			User: models.User{
				ID:    1,
				Name:  "test user",
				Image: "user.png",
			},
			Reactions: "{}",
		},
	}

	// モック関数に差し替え
	models.GetMessages = func(channelID string) ([]models.Message, error) {
		return mockMessages, nil
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages?channel_id=test123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ListMessages(c)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `[{"id":1,"content":"test message","channel_id":1,"created_at":"2025-04-03T21:58:47.107456+09:00","user":{"id":1,"name":"test user","image":"user.png","email":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"reactions":"{}"}]`, rec.Body.String())
}

// todo AddMessageの分を書く

// まずリクエスト失敗時
func TestAddMessage_InvalidRequest(t *testing.T) {
	e := echo.New()
	// 無効なJSONを送信
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader("invalid json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "リクエストの形式が正しくありません")
}

// メッセージが空のとき
func TestAddMessage_EmptyMessage(t *testing.T) {
	e := echo.New()
	reqBody := `{"content":"","channel_id":1,"user":{"id":1}}`
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "Messageが空です")
}

// チャンネルIDが空のとき
func TestAddMessage_EmptyChannelID(t *testing.T) {
	e := echo.New()
	reqBody := `{"content":"test message","channel_id":0,"user":{"id":1}}`
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "ChannelIDが必要です")
}

// UserIDが空のとき
func TestAddMessage_EmptyUserID(t *testing.T) {
	e := echo.New()
	reqBody := `{"content":"test message","channel_id":1,"user":{"id":0}}`
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "UserIDが必要です")
}

// DB insert時にエラーが出た時
func TestAddMessage_DBError(t *testing.T) {
	original := models.AddMessage
	defer func() { models.AddMessage = original }()

	// モック関数に差し替え
	models.AddMessage = func(content string, channelID int, user models.User) (models.Message, error) {
		return models.Message{}, errors.New("DB connection failed")
	}

	e := echo.New()
	reqBody := `{"content":"test message","channel_id":1,"user":{"id":1}}`
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.Contains(t, rec.Body.String(), "データベースエラー: DB connection failed")
}

// 作成成功時に返ってくるmessageが正しい
func TestAddMessage_Success(t *testing.T) {
	e := echo.New()

	original := models.AddMessage
	defer func() { models.AddMessage = original }()

	models.AddMessage = func(content string, channelID int, user models.User) (models.Message, error) {
		return models.Message{
			ID:        1,
			Content:   "test message",
			CreatedAt: time.Date(2025, 4, 3, 21, 58, 47, 107456000, time.FixedZone("JST", 9*60*60)),
			ChannelID: channelID,
			User: models.User{
				ID:    1,
				Name:  "test user",
				Image: "user.png",
			},
			Reactions: "{}",
		}, nil
	}

	reqBody := `{"content":"test message","channel_id":1,"user":{"id":1,"name":"test user","image":"user.png"}}`
	req := httptest.NewRequest(http.MethodPost, "/messages", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := AddMessage(c)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, rec.Code)
	require.JSONEq(t, `{"id":1,"content":"test message","channel_id":1,"created_at":"2025-04-03T21:58:47.107456+09:00","user":{"id":1,"name":"test user","image":"user.png","email":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"reactions":"{}"}`, rec.Body.String())
}