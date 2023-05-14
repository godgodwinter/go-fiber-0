package responses

import "github.com/gofiber/fiber/v2"

type DefaultResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}