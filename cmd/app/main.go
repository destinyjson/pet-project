package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/messageService"
	"pet-project/internal/userService"
	"pet-project/internal/web/messages"
	"pet-project/internal/web/users"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messageService.RequestBody{}, &userService.User{})
	if err != nil {
		return
	}

	repoMsg := messageService.NewMessageRepository(database.DB)
	serviceMsg := messageService.NewService(repoMsg)
	msgHandler := handlers.NewMessageHandler(serviceMsg)

	repoUsr := userService.NewUserRepository(database.DB)
	serviceUsr := userService.NewService(repoUsr)
	usrHandler := handlers.NewUserHandler(serviceUsr)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictMsgHandler := messages.NewStrictHandler(msgHandler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictMsgHandler)

	strictUsrHandler := users.NewStrictHandler(usrHandler, nil)
	users.RegisterHandlers(e, strictUsrHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
