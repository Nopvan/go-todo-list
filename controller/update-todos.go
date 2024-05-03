package controller

import (
	"database/sql"
	"encoding/json"
	"go-todo-list/models"
	"net/http"

	"github.com/labstack/echo"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateTodos(e *echo.Echo, db *sql.DB) {

	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		//untuk input user_id di todos
		user := ctx.Get("USER").(models.AuthClaimJWT)

		//set permission untuk user
		permissionFound := false
		for _, scope := range user.UserScopes {
			if scope == "todos:update" {
				permissionFound = true
				break
			}
		}

		if !permissionFound {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		var request UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ? AND user_id = ?",
			request.Title,
			request.Description,
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
