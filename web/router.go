package web

import (
	"log"
	"net/http"

	"github.com/bitxeno/go-docker-skeleton/internal/app"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func route(fi *fiber.App) {
	fi.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(StaticAssets()),
	}))
	fi.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello world.")
	})

	// websocket router
	fi.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	fi.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Params("id")) // 123
		log.Println(c.Query("v"))   // 1.0

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	// api router
	api := fi.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello world.")
	})
	api.Get("/setting", func(c *fiber.Ctx) error {
		return c.JSON(apiSuccess(app.Settings))
	})
	api.Post("/setting", func(c *fiber.Ctx) error {
		var setting app.SettingsConfiguration
		if err := c.BodyParser(&setting); err != nil {
			return c.JSON(apiError("Invalid argument. error: " + err.Error()))
		}

		// update settings
		app.Settings.Test = setting.Test
		app.SaveSettings()

		return c.JSON(apiSuccess(true))
	})
}
