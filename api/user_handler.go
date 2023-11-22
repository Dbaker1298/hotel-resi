package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "David Baker"})
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Michael Baker"})
}
