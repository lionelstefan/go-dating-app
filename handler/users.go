package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lionelstefan/go-dating-app/model/request"
	"github.com/lionelstefan/go-dating-app/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lionelstefan/go-dating-app/config"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}
	
	user, err := repository.FindByCredentials(loginRequest.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}

	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if match != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong Password !",
		   })
	}

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":  user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Register(c *fiber.Ctx) error {
	registerRequest := new(request.RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}

	registerModel := registerRequest.ToModel()
	res, err := repository.RegisterNewUser(registerModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}

	log.Infof("Register Success", res)
	return c.JSON(fiber.Map{"success": "true"})
}