package main

import (
	"goHexagonalBlog/internal/core/services"
	"goHexagonalBlog/internal/handlers"
	"goHexagonalBlog/internal/repositories"
	"goHexagonalBlog/internal/server"
)

func main() {
	mongoConn := "secret🤫"
	//repositories
	userRepository := repositories.NewUserRepository(mongoConn)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//server
	server.Initialize(
		userHandlers,
	)
}
