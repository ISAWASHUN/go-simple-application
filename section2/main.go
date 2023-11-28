package main

import (
	"database/sql"
	"log"
	"section2/controller"
	"section2/repository"
	"section2/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func intiDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {

	db, err := intiDB()
	if err != nil {
		log.Fatal()
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT)")
	if err != nil {
		log.Fatal()
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskRepository := repository.NewTaskRepository(db)
	taskusecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskusecase)

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}