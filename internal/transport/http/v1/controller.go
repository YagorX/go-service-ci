package v1

import (
	"context"

	"github.com/YagorX/go-service-ci/internal/model"
	"github.com/labstack/echo/v4"
)

type SatelliteService interface {
	GetSatelliteByName(ctx context.Context, name string) (*model.Satellite, error)
}

type Controller struct {
	service SatelliteService
}

func NewController(g *echo.Group, s SatelliteService) {
	ctl := &Controller{
		service: s,
	}

	g.GET("/:name", ctl.Get)
}
