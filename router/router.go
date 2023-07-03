package router

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/bitxeno/go-docker-skeleton/config"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func Create(app *fiber.App, f fs.FS) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(f),
	}))
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello world.")
	})

	// websocket router
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
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
	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello world.")
	})
	api.Get("/setting", func(c *fiber.Ctx) error {
		return c.JSON(apiSuccess(config.Settings))
	})
	api.Post("/setting", func(c *fiber.Ctx) error {
		var setting config.SettingsConfiguration
		if err := c.BodyParser(&setting); err != nil {
			return c.JSON(apiError("Invalid argument. error: " + err.Error()))
		}

		config.SaveSettings()
		return c.JSON(apiSuccess(true))
	})
}
