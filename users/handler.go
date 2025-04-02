package users

import (
	"fmt"
	"subscription_tracker/users/db"

	"github.com/labstack/echo"
)

func (h *Handler) CreateUser(c echo.Context) error {
	fmt.Println("Creating a new user")
	return c.JSON(200, "user")
}

// this function gets all users
func (h *Handler) GetUsers(c echo.Context) error {
	fmt.Println("Getting all users")
	usersModel := &db.Users{}
	users, err := usersModel.GetAllUsers(c.Request().Context(), h.app)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, users)

}
