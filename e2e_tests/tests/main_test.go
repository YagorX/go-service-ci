//go:build test_e2e

package test

import (
	"os"
	"testing"
)

var appURL = "http://app:10080"

func TestMain(m *testing.M) {
	if value, ok := os.LookupEnv("APP_URL"); ok && value != "" {
		appURL = value
	}

	os.Exit(m.Run())
}
