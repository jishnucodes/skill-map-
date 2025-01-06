package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jishnucodes/skill-map/common"
	"github.com/jishnucodes/skill-map/managers"
)

type AuthHandler struct {
	groupName   string
	userManager managers.UserManager
	secretKey   string
}

func NewAuthHandler(userManager managers.UserManager, secretKey string) *AuthHandler {
	return &AuthHandler{
		"api/auth",
		userManager,
		secretKey,
	}
}

func (handler *AuthHandler) RegisterAuthApis(r *gin.Engine) {
	authGroup := r.Group(handler.groupName);
	authGroup.POST("/signup", handler.SignUp)
	authGroup.POST("/signin", handler.SignIn)
}

func (handler *AuthHandler) SignUp(ctx *gin.Context) {
	fmt.Println("Signup")

	// Get the user data from the request
	userData := common.NewUserCreationInput()
	err := ctx.BindJSON(&userData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userExist, _ := handler.userManager.FindSingleUserByEmail(userData.Email)
	if userExist != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//password hashing
	hashedPassword, err := common.HashPassword(userData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	userData.Password = hashedPassword

	// Create the user
	newUser, err := handler.userManager.CreateUser(userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	newUser.Password = "" // Remove the password from the response
	// Return the user data
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"data": newUser,
	})


}

func (handler *AuthHandler) SignIn(ctx *gin.Context) {
	fmt.Println("Signin")

	userData := common.NewUserSigninInput()
    err := ctx.BindJSON(&userData)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

	existingUser, _:= handler.userManager.FindSingleUserByEmail(userData.Email)

    
    if existingUser == nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

	// Check if the password is correct
	isValidPassword := common.CheckPasswordHash(userData.Password, existingUser.Password)
	if !isValidPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate a JWT token
	token, _ := common.GenerateJWT(existingUser.ID, handler.secretKey)
	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.SetCookie("token", token, 3600, "/", "", true, false)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user signed in successfully",
		"data": existingUser,
	})



	
}

