package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jishnucodes/skill-map/database"
	"github.com/jishnucodes/skill-map/handlers"
	"github.com/jishnucodes/skill-map/managers"
	
)

func init() {
	database.Initialize()
}

func main() {
	fmt.Println("Hello, World!")


	r := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandler(userManager)
	userHandler.RegisterUserApis(r)

	r.Run() 
}
