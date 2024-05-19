package handler

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/lionelstefan/go-dating-app/model"
	"github.com/lionelstefan/go-dating-app/model/request"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lionelstefan/go-dating-app/repository"
	"github.com/google/uuid"
)

func Swipe(c *fiber.Ctx) error {
	swipeRequest := new(request.SwipeRequest)
	if err := c.BodyParser(swipeRequest); err != nil {
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

	targetId, err := uuid.Parse(swipeRequest.TargetId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	isLike, err := strconv.ParseBool(swipeRequest.IsLike)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	mdl := model.Swipes{
		UserId: userId,
		TargetId: targetId,
		IsLike: isLike,
	}

	err = repository.AddNewSwipe(mdl)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		 "error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"success": true})
}