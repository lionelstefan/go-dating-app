package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/lionelstefan/go-dating-app/router"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New())
	
	router.Initialize(app)
	app.Listen(":3000")
}