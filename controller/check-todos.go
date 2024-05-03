package controller

import (
	"database/sql"
	"encoding/json"
	"go-todo-list/models"
	"net/http"

	"github.com/labstack/echo"
)

type CheckRequest struct {
	Done bool `json:"done"`
}

func CheckTodos(e *echo.Echo, db *sql.DB) {

	e.PUT("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")
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

		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneInt int
		if request.Done {
			doneInt = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ? AND user_id = ?",
			doneInt,
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
