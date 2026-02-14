package application

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BigDwarf/testci/internal/model"
	"github.com/BigDwarf/testci/internal/repository/satellite"
	"github.com/BigDwarf/testci/internal/service/cache"
	"github.com/redis/go-redis/v9"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/BigDwarf/testci/internal/config"
	satelliteService "github.com/BigDwarf/testci/internal/service/satellite"
	v1 "github.com/BigDwarf/testci/internal/transport/http/v1"
)

type App struct {
	cfg *config.Config

	db          *sql.DB
	redisClient *redis.Client
	srv         *echo.Echo
}

func NewApp() *App {
	srv := echo.New()

	return &App{
		srv: srv,
		cfg: config.NewDefaultConfig(),
	}
}

func (app *App) Start() error {
	app.RegisterRoutes()

	go func() {
		if err := app.srv.Start(app.cfg.HttpServerConfig.ListenAddress); err != nil {
			fmt.Println("error not null, ", err.Error())
			log.Info(err)
		}
	}()

	return nil
}

func (app *App) Stop(ctx context.Context) {
	err := app.srv.Shutdown(ctx)
	if err != nil {
		fmt.Println("error not null, ", err.Error())
		log.Info(err)
	}
}

func (app *App) RegisterRoutes() {
	g := app.srv.Group("/api/v1/satellite")
	db := app.Database(app.cfg.Database.GetDSN())
	app.redisClient = redis.NewClient(&redis.Options{Addr: app.cfg.RedisAddr})
	satelliteCache := cache.New[*model.Satellite](app.redisClient)

	v1.NewController(g, satelliteService.NewService(satellite.NewRepository(db), satelliteCache))
	app.registerHealthRoutes()
}
