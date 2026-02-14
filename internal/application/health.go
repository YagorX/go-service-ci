package application

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type healthResource struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type healthResponse struct {
	Status    string                    `json:"status"`
	Resources map[string]healthResource `json:"resources"`
}

func (app *App) registerHealthRoutes() {
	app.srv.GET("/health/check", app.healthCheck)
}

func (app *App) healthCheck(ctx echo.Context) error {
	resp := healthResponse{
		Status: "ok",
		Resources: map[string]healthResource{
			"database": {Status: "ok"},
			"redis":    {Status: "ok"},
		},
	}

	checkCtx, cancel := context.WithTimeout(ctx.Request().Context(), 2*time.Second)
	defer cancel()

	if err := app.db.PingContext(checkCtx); err != nil {
		resp.Status = "error"
		resp.Resources["database"] = healthResource{Status: "error", Error: err.Error()}
	}

	if err := app.redisClient.Ping(checkCtx).Err(); err != nil {
		resp.Status = "error"
		resp.Resources["redis"] = healthResource{Status: "error", Error: err.Error()}
	}

	return ctx.JSON(http.StatusOK, resp)
}
