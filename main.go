package main

import (
	"go-todo-list/controller"
	"go-todo-list/database"

	"go-todo-list/middleware"

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

	e.Use(middleware.AuthMiddleware)

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

	//CREATE USERS
	controller.Register(e, db)

	//LOGIN USERS
	controller.Login(e, db)

	//CREATE SCOPES
	controller.CreateScopes(e, db)

	//DELETE SCOPES
	controller.DeleteScopes(e, db)

	//ASSIGN SCOPES TO USER
	controller.AssignScope(e, db)

	//RUN IN PORT 8080
	e.Start(":8080")
}
