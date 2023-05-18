package services

import (
	"context"
	"time"

	"github.com/godgodwinter/go-fiber-0/app/types"
	"github.com/gofiber/fiber/v2"
)

type UserService struct{}

func (s *UserService) GetUser(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	data := types.ExampleStruct{
		Name: c.Params("userId"),
		Desc: "Ini userId " + userId,
	}
	return c.JSON(data)
}
