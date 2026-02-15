//go:build test_e2e

package test

import (
	"net/http"
	"testing"
)

func TestGetSatellite(t *testing.T) {
	req, err := http.NewRequest("GET", appURL+"/api/v1/satellite/moon", nil)
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
}
