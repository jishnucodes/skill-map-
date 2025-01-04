package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jishnucodes/skill-map/common"
	"github.com/jishnucodes/skill-map/managers"
)

type PostHandler struct {
	groupName   string
	postManager *managers.PostManager
}

func NewPostHandler(postManager *managers.PostManager) *PostHandler {
	return &PostHandler{
		"api/posts",
		postManager,
	}

}

func (handler *PostHandler) RegisterPostApis(r *gin.Engine) {
	postGroup := r.Group(handler.groupName)
	postGroup.GET("", handler.GetPosts)
	postGroup.POST("/create", handler.CreatePost)
}

func (handler *PostHandler) GetPosts(ctx *gin.Context)  {

	posts, err := handler.postManager.GetPosts()
	if err!= nil {
		 ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
	  "message": "success",
	  "data": posts,
	})

}

func (handler *PostHandler) CreatePost(ctx *gin.Context)  {

	postData := common.NewPostCreationInput()

	err :=ctx.BindJSON(&postData)

	if err!= nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost, err := handler.postManager.CreatePost(postData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": newPost,
	})

}