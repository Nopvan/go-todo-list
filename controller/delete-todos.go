package controller

import (
	"database/sql"
	"go-todo-list/models"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteTodos(e *echo.Echo, db *sql.DB) {

	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		user := ctx.Get("USER").(models.AuthClaimJWT)
		id := ctx.Param("id")

		_, err := db.Exec(
			"DELETE FROM todos WHERE id = ? AND user_id = ?",
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
