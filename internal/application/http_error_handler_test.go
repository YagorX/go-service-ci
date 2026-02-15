package application

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHTTPErrorHandler_CommittedResponse(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/unknown", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Response().Committed = true

	httpErrorHandler(echo.NewHTTPError(http.StatusNotFound), c)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusOK)
	}
	if rec.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %q", rec.Body.String())
	}
}

func TestHTTPErrorHandler_NotFoundHTML(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/some-page", nil)
	req.Header.Set("Accept", "text/html")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	httpErrorHandler(echo.NewHTTPError(http.StatusNotFound), c)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusNotFound)
	}
	if !strings.Contains(rec.Body.String(), "<!doctype html>") {
		t.Fatalf("expected html response, got %q", rec.Body.String())
	}
}

func TestHTTPErrorHandler_NotFoundJSONForAPI(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/unknown", nil)
	req.Header.Set("Accept", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	httpErrorHandler(echo.NewHTTPError(http.StatusNotFound), c)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusNotFound)
	}
	if !strings.Contains(rec.Body.String(), `"error":"resource not found"`) {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHTTPErrorHandler_InternalServerError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/satellite/moon", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	httpErrorHandler(errors.New("boom"), c)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("unexpected status code: got %d want %d", rec.Code, http.StatusInternalServerError)
	}
	if !strings.Contains(rec.Body.String(), `"error":"Internal Server Error"`) {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
