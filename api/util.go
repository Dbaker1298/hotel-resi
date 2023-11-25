package api

import (
	"fmt"

	"github.com/Dbaker1298/hotel-resi/types"

	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("not authorized")
	}

	return user, nil
}