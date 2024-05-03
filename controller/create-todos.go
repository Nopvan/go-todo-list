package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-todo-list/models"
	"net/http"

	"github.com/labstack/echo"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTodos(e *echo.Echo, db *sql.DB) {

	e.POST("/todos", func(ctx echo.Context) error {

		//untuk input user_id di todos
		user := ctx.Get("USER").(models.AuthClaimJWT)

		fmt.Println(user.UserScopes)

		//set permission untuk user
		permissionFound := false
		for _, scope := range user.UserScopes {
			if scope == "todos:create" {
				permissionFound = true
				break
			}
		}

		if !permissionFound {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done, user_id) VALUES (?, ?, 0, ?)",
			request.Title,
			request.Description,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")

	})
}
