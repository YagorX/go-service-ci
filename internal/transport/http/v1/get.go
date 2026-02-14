package v1

import (
	"errors"
	"net/http"

	satelliterepo "github.com/YagorX/go-service-ci/internal/repository/satellite"

	"github.com/labstack/echo/v4"
)

func (ctl *Controller) Get(ctx echo.Context) error {
	name := ctx.Param("name")

	satellite, err := ctl.service.GetSatelliteByName(ctx.Request().Context(), name)
	if err != nil {
		if errors.Is(err, satelliterepo.ErrSatelliteNotFound) {
			return ctx.JSON(http.StatusNotFound, err.Error())
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, satellite)
}
