//go:build test_e2e

package test

import (
	"encoding/json"
	"net/http"
	"testing"
)

type healthCheckResponse struct {
	Status    string                         `json:"status"`
	Resources map[string]healthCheckResource `json:"resources"`
}

type healthCheckResource struct {
	Status string `json:"status"`
}

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", appURL+"/health/check", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d", resp.StatusCode, http.StatusOK)
	}

	var body healthCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode health response: %v", err)
	}

	if body.Status != "ok" {
		t.Fatalf("unexpected health status: got %q want %q", body.Status, "ok")
	}

	for _, key := range []string{"database", "redis"} {
		resource, ok := body.Resources[key]
		if !ok {
			t.Fatalf("missing health resource: %s", key)
		}
		if resource.Status != "ok" {
			t.Fatalf("resource %s status: got %q want %q", key, resource.Status, "ok")
		}
	}
}
