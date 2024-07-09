package services

import (
	"fmt"

	"github.com/hazitgi/graphql-with-fiber/apis/common"
	"github.com/hazitgi/graphql-with-fiber/config"
	"github.com/hazitgi/graphql-with-fiber/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService() *UserService {
	db := config.GetDb()
	return &UserService{DB: db}
}

func (userSrv *UserService) InsertUser(user *models.User) (*models.User, error) {
	result := userSrv.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userSrv *UserService) GetUserByID(id uint) (*models.User, error) {
	user := models.User{}
	result := userSrv.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userSrv *UserService) UpdateUser(id uint, updatedUser common.CreateUserInput) (*models.User, error) {
	user := models.User{}
	result := userSrv.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	user.FullName = updatedUser.FullName
	user.CompanyName = updatedUser.CompanyName
	user.CountryID = updatedUser.CountryID
	user.StateID = updatedUser.StateID
	user.Email = updatedUser.Email
	user.Location = updatedUser.Location
	user.Address = updatedUser.Address

	result = userSrv.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userSrv *UserService) DeleteUser(id uint) error {
	result := userSrv.DB.Unscoped().Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (userSrv *UserService) FindAllUsers(pagination *common.Pagination) ([]*models.User, error) {
	var users []*models.User
	// Set default values if not provided
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Sort == "" {
		pagination.Sort = "asc"
	}
	pagination.Offset = (pagination.Page - 1) * pagination.Limit

	result := userSrv.DB.Order(fmt.Sprintf("%s ", pagination.SortField) + pagination.Sort).Limit(pagination.Limit).Offset(pagination.Offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (userSrv *UserService) FindUserByField(field, value string) (*models.User, error) {
	fmt.Printf("Finding user by %s '%s'... \n", field, value)
	user := models.User{}
	// query := userSrv.DB.Model(&models.User{})

	result := userSrv.DB.Where(fmt.Sprintf("%s =?", field), value).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
