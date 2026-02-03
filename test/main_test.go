//go:build test_integration

package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/BigDwarf/testci/internal/application"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	os.Exit(m.Run())
}

func createAndRunApp(_ context.Context) *application.App {
	app := application.NewApp()
	app.Start()
	return app
}
