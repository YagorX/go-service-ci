package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctl *Controller) Get(ctx echo.Context) error {
	name := ctx.Param("name")

	satellite, err := ctl.service.GetSatelliteByName(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, satellite)
}
