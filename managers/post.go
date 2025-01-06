package managers

import (
	// "errors"

	"errors"
	"fmt"

	"github.com/jishnucodes/skill-map/common"
	"github.com/jishnucodes/skill-map/database"
	"github.com/jishnucodes/skill-map/models"
)

type PostManager struct {
}

func NewPostManager() *PostManager {
	return &PostManager{}
}

func (pm *PostManager) GetPosts() ([]models.Post, error) {
	posts := []models.Post{}

	result := database.DB.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil

}

func (pm *PostManager) CreatePost(postData *common.PostCreationInput) (*models.Post,error) {
	newPost := &models.Post{Title: postData.Title, Content: postData.Content, Author: postData.Author}
	database.DB.Create(newPost)

	if newPost.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newPost, nil

}

func (pm *PostManager) GetSinglePost(id string) (*models.Post, error) {
	post := models.Post{}

	result := database.DB.First(&post, id)

	if result.Error != nil {
		fmt.Println("Error in getting post", result.Error)
		return nil, result.Error
	}

	return &post, nil

}

func (pm *PostManager) DeletePost(id string) (error) {
	post := models.Post{}

	result := database.DB.Delete(&post, id)

	if result.Error != nil {
		fmt.Println("Error in getting post", result.Error)
		return result.Error
	}

	return  nil

}

func (pm *PostManager) UpdatePost(id string, postData *common.PostUpdateInput) (*models.Post,error) {

	post := models.Post{}
	result := database.DB.First(&post, id)

	if result.Error != nil {
		fmt.Println("Error in getting post", result.Error)
		return nil, result.Error
	}
	database.DB.Model(&post).Updates(&models.Post{Title: postData.Title, Content: postData.Content, Author: postData.Author})

	return &post, nil

}
