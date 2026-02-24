package main

import (
	"net/http"

	"github.com/hirakan045/feelog/backend/internal/handler"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

    feelogHandler := handler.FeelogHandler{}

    	// ルーティング
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Feelog API is running!",
		})
	})

    api := e.Group("/api/v1")
	api.POST("/feelogs", feelogHandler.CreateFeelog)
	api.GET("/feelogs", feelogHandler.GetFeelogs)

	e.Start(":8080")
}