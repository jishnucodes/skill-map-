package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jishnucodes/skill-map/common"
	"github.com/jishnucodes/skill-map/managers"
)

type UserHandler struct {
	groupName string
	userManager managers.UserManager
}

func NewUserHandler(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		 "api/users",
		 userManager,
	}
}

func (handler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)
	userGroup.GET("", handler.List)
	userGroup.POST("", handler.Create)
	userGroup.GET("/:user_id", handler.Detail)
	userGroup.DELETE("/:user_id", handler.Delete)
	userGroup.PUT("/:user_id", handler.Update)


}


	

func (handler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)


	if err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := handler.userManager.CreateUser(userData)

	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	common.SuccessResponse(ctx, "User created successfully", newUser)

	
}	

func (handler *UserHandler) List(ctx *gin.Context) {

	
	users, err := handler.userManager.UsersList()

	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Listed successfully",
		"data":    users,
	})
}

func (handler *UserHandler) Detail(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("user_id")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is missing"})
		return
	}

	
	user, err := handler.userManager.UserDetail(userId)

	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Listed successfully",
		"data":    user,
	})
}

func (handler *UserHandler) Delete(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("user_id")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is missing"})
		return
	}

	
	err := handler.userManager.DeleteUser(userId)

	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// common.SuccessResponse(ctx, "User deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("user_id")

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is missing"})
		return
	}

	userData := common.NewUserUpdateInput()

	err := ctx.BindJSON(&userData)


	if err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := handler.userManager.UpdateUser(userId,userData)

	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	common.SuccessResponse(ctx, "User created successfully", updatedUser)

	
}	