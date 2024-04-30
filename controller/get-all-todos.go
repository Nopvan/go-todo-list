package controller

import (
	"database/sql"
	"go-todo-list/models"
	"net/http"

	"github.com/labstack/echo"
)

type TodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func GetAllTodos(e *echo.Echo, db *sql.DB) {

	e.GET("/todos", func(ctx echo.Context) error {

		//untuk set hanya user id yang cocok saja yang di tampilkan
		user := ctx.Get("USER").(models.AuthClaimJWT)

		rows, err := db.Query("SELECT id, title, description, done FROM todos WHERE user_id = ?", user.UserId)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []TodoResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int

			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}

			res = append(res, todo)
		}

		if len(res) == 0 {
			return ctx.JSON(http.StatusOK, map[string]string{"message": "to do not found"})
		}

		return ctx.JSON(http.StatusOK, res)
	})
}
