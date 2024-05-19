package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/lionelstefan/go-dating-app/handler"
	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/lionelstefan/go-dating-app/config"
)

func Initialize(app *fiber.App) error {
	jwt := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.Secret),
		},
	})

	app.Post("/api/login", handler.Login)
	
	app.Post("/api/register", handler.Register)
	
	app.Post("/api/swipe", jwt, handler.Swipe)

	app.Post("/api/purchase/premium", jwt, handler.PurchasePremium)

	return nil
}