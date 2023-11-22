package api

import (
	"github.com/Dbaker1298/hotel-resi/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "David",
		LastName:  "Baker",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Michael Baker"})
}
