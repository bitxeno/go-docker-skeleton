package main

import (
	"fmt"
	"os"

	"github.com/bitxeno/go-docker-skeleton/config"
	"github.com/bitxeno/go-docker-skeleton/internal/app"
	"github.com/bitxeno/go-docker-skeleton/internal/db"
	_ "github.com/bitxeno/go-docker-skeleton/internal/log"
	"github.com/bitxeno/go-docker-skeleton/model"
	"github.com/bitxeno/go-docker-skeleton/router"
	"github.com/gofiber/fiber/v2"
)

const (
	// service name
	AppName = "go-docker-skeleton"
	// service description
	AppDesc = "Skeleton for run go service in docker"
)

func main() {
	app := app.New(AppName, AppDesc)
	app.Route(func(f *fiber.App) {
		router.Create(f, ViewAssets())
	})
	app.AddBoot(func() error {
		// init config
		if err := config.Load(); err != nil {
			return err
		}

		// init db
		return db.Open().AutoMigrate(&model.User{})
	})

	if err := app.Run(os.Args); err != nil {
		code := 1
		fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}
}
