package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(e *echo.Echo, db *sql.DB) {

	e.POST("/auth/register", func(ctx echo.Context) error {

		var request RegisterRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		bytePassword := []byte(request.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, 10)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		_, err = db.Exec(
			"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
			request.Name,
			request.Email,
			string(hashedPassword),
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")

	})
}
