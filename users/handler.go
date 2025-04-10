package users

import (
	"fmt"
	"subscription_tracker/users/db"

	"github.com/labstack/echo"
)

func (h *Handler) CreateUser(c echo.Context) error {
	userData, err := validateCreateUserData(c)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	usersModel := &db.Users{}
	newUser, err := usersModel.CreateUser(c.Request().Context(), h.app, userData)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, newUser)
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
