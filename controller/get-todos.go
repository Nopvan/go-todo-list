package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type OneTodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func GetTodos(e *echo.Echo, db *sql.DB) {
	e.GET("/todos/:id", func(ctx echo.Context) error {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid ID")
		}

		var response OneTodoResponse
		err = db.QueryRow("SELECT id, title, description, done FROM todos WHERE id = ?", id).
			Scan(&response.Id, &response.Title, &response.Description, &response.Done)
		if err != nil {
			if err == sql.ErrNoRows {
				return ctx.String(http.StatusNotFound, "Todo not found")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, response)
	})
}
