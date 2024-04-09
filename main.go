package main

import (
	"go-todo-list/controller"
	"go-todo-list/database"

	"github.com/labstack/echo"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	//CREATE TODOS
	controller.CreateTodos(e, db)

	//READ TODOS
	controller.GetAllTodos(e, db)

	//UPDATE TODOS
	controller.UpdateTodos(e, db)

	//DELETE TODOS
	controller.DeleteTodos(e, db)

	//UPDATE CHECK
	controller.CheckTodos(e, db)

	//GET TODOS
	controller.GetTodos(e, db)

	e.Start(":8080")
}
