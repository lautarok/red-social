package infra

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lautarok/yorcom/internal/delivery/api"
	"github.com/uptrace/bun"
)

func InitApi(db *bun.DB) {
	router := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	router.Hooks().OnListen(func(ld fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}

		scheme := "http"
		if ld.TLS {
			scheme = "https"
		}

		log.Println("Server started on " + scheme + "://" + ld.Host + ":" + ld.Port)

		return nil
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
	}))

	v1_router := router.Group("/api/v1/")

	api.InitRoutes(v1_router, db)

	log.Fatal(
		router.Listen(os.Getenv("ADDRESS")),
	)
}
