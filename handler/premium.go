package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lionelstefan/go-dating-app/model"
	"github.com/lionelstefan/go-dating-app/model/request"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lionelstefan/go-dating-app/repository"
	"github.com/google/uuid"
)

func PurchasePremium (c *fiber.Ctx) error {
	premiumRequest := new(request.PremiumRequest)
	if err := c.BodyParser(premiumRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	userId, err := uuid.Parse(user_id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	mdl := model.PremiumPurchase{
		UserId: userId,
		PremiumType: premiumRequest.PremiumType,
	}

	err = repository.Purchase(mdl)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true})
}