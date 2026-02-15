package application

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const notFoundPageHTML = `<!doctype html>
<html lang="ru"><head><meta charset="utf-8"><title>404</title></head>
<body style="font-family:Segoe UI,Tahoma,sans-serif;padding:40px;background:#f6f8ff;color:#1a2340">
  <h1>404: Страница пока не существует</h1>
  <p>Мы уже работаем над этим разделом.</p>
  <p>Попробуйте вернуться позже.</p>
  <a href="/">На главную</a>
</body></html>`

func httpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == http.StatusNotFound {
		accept := c.Request().Header.Get("Accept")
		if strings.Contains(accept, "text/html") || !strings.HasPrefix(c.Request().URL.Path, "/api") {
			_ = c.HTML(http.StatusNotFound, notFoundPageHTML)
			return
		}
		_ = c.JSON(http.StatusNotFound, map[string]string{
			"error": "resource not found",
		})
		return
	}

	_ = c.JSON(code, map[string]string{"error": http.StatusText(code)})
}
