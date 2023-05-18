// user_controller.go

package controllers

import (
	"github.com/godgodwinter/go-fiber-0/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetAUser(c *fiber.Ctx) error {
	service := services.UserService{}
	return service.GetUser(c)
}

// dengan middleware
func GetAUserMd(c *fiber.Ctx) error {
	service := services.UserService{}
	return service.GetUser(c)
}
