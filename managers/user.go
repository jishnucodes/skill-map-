package managers

import (
	"errors"

	"github.com/jishnucodes/skill-map/common"
	"github.com/jishnucodes/skill-map/database"
	"github.com/jishnucodes/skill-map/models"
)
type UserManager interface {
	CreateUser(userData *common.UserCreationInput) (*models.User, error)
	UsersList() ([]models.User, error)
	UserDetail(id string) (*models.User, error)
	DeleteUser(id string) (error)
	UpdateUser(id string, userData *common.UserUpdateInput) (*models.User, error)
}

type userManager struct {
	// dbClient
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (um *userManager) CreateUser(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{Name: userData.Name, Email: userData.Email, Password: userData.Password}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newUser, nil
}

func (um *userManager) UsersList() ([]models.User, error) {

	users := []models.User{}
	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (um *userManager) UserDetail(id string) (*models.User, error) {
	user := models.User{}
	result := database.DB.First(&user, id)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}


func (um *userManager) DeleteUser(id string) (error) {
	user := models.User{}
	result := database.DB.Delete(&user, id)

	if result.Error != nil {
		return errors.New("user not found")
	}

	return nil
}


func (um *userManager) UpdateUser(id string, userData *common.UserUpdateInput) (*models.User, error) {

	user := models.User{}
	result := database.DB.First(&user, id)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	// user.Name = userData.Name
	// user.Email = userData.Email
	// user.Password = userData.Password

	// database.DB.Save(&user)

	database.DB.Model(&user).Updates(&models.User{Name: userData.Name, Email: userData.Email, Password: userData.Password})

	

	if user.ID == 0 {
		return nil, errors.New("user updation failed")
	}

	return &user, nil
}