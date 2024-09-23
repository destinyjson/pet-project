package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/service"
	"pet-project/internal/web/messages"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&service.RequestBody{})
	if err != nil {
		return
	}

	repo := service.NewMessageRepository(database.DB)
	serviceMsg := service.NewService(repo)

	handler := handlers.NewHandler(serviceMsg)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := messages.NewStrictHandler(handler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
