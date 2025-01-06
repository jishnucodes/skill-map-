package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jishnucodes/skill-map/database"
	"github.com/jishnucodes/skill-map/handlers"
	"github.com/jishnucodes/skill-map/managers"
	"github.com/joho/godotenv"
)

func init() {
	database.Initialize()
}

func main() {
	fmt.Println("Hello, World!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET")

	r := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandler(userManager)
	userHandler.RegisterUserApis(r)

	postManager := managers.NewPostManager()
	postHandler := handlers.NewPostHandler(postManager)
	postHandler.RegisterPostApis(r)

	authManager := managers.NewUserManager()
	authHandler := handlers.NewAuthHandler(authManager, secretKey)
	authHandler.RegisterAuthApis(r)

	r.Run()
}
