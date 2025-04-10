package users

import (
	"fmt"
	"subscription_tracker/users/db"

	"github.com/labstack/echo"
)

func validateCreateUserData(c echo.Context) (db.Users, error) {
	name := c.FormValue("name")
	if name == "" {
		fmt.Println("Error: Missing first name")
		return db.Users{}, fmt.Errorf("missing first name")
	}
	email := c.FormValue("email")
	if email == "" {
		fmt.Println("Error: Missing email")
		return db.Users{}, fmt.Errorf("missing email")
	}
	password := c.FormValue("password")
	if password == "" {
		fmt.Println("Error: Missing password")
		return db.Users{}, fmt.Errorf("missing password")
	}

	return db.Users{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
