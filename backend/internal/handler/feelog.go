package handler

import (
	"net/http"
	"time"

	"github.com/hirakan045/feelog/backend/internal/model"
	"github.com/labstack/echo/v5"
)

// メモリ上でログを管理するスライス
var feelogs = make([]model.Feelog, 0)	// 初期化: 空のスライスを作成
var nextID = 1

type FeelogHandler struct{}

// POST /api/v1/feelogs
func (h *FeelogHandler) CreateFeelog(c *echo.Context) error {
	log := new(model.Feelog)

	if err := c.Bind(log); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	// バリデーション
	if log.FeelScore < 1 || log.FeelScore > 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "feel score must be between 1 and 5",
		})
	}

	log.ID = nextID
	log.CreatedAt = time.Now()
	nextID++

	feelogs = append(feelogs, *log)

	return c.JSON(http.StatusCreated, log)
}

// GET /api/v1/feelogs
func (h *FeelogHandler) GetFeelogs(c *echo.Context) error {
	return c.JSON(http.StatusOK, feelogs)
}