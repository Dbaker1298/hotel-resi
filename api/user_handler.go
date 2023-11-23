package api

import (
	"context"

	"github.com/Dbaker1298/hotel-resi/db"
	"github.com/Dbaker1298/hotel-resi/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// Not the specifice store, but the interface
	userStore db.UserStore
}

// constructor; a factory function
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "David",
		LastName:  "Baker",
	}
	return c.JSON(u)
}
