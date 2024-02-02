package main

import (
    "golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/labstack/echo/v4"
	"strings"
)

type User struct {
	Username string
	Password string
}

var users = map[string]User{
	"Nisse": {Username: "Nisse", Password: "Nisse900"},
}

func main() {
	e := echo.New()

	e.GET("/login", func(c echo.Context) error {
		return c.File("login.html")
	})

	e.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

enteredUsernameLower := strings.ToLower(username)

user, ok := users[enteredUsernameLower]
if !ok {
    return c.String(http.StatusUnauthorized, "Invalid log-in details")
}

	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid log-in details")
	}
	return c.String(http.StatusOK, "Login successful")
})
e.Start(":8080")
}

	
