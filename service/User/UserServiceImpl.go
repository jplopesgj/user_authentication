package services

import (
	database "finance_backend/database/config"
	models "finance_backend/models/User"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	c *gin.Context
}

// func NewUserServiceImpl() UserService {
// 	return &UserServiceImpl{}
// }

func (us *UserServiceImpl) GetSpecificUser(c *gin.Context) {
	id := c.Param("id")
	user, err := checkIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (us *UserServiceImpl) RegisterUser(c *gin.Context) {
	var newUser models.CreateUserFields

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := newUser.HashPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := userExist(newUser)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email Already exist"})
		c.Abort()
		return
	}

	record := database.DB.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (us *UserServiceImpl) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := checkIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var userInput models.UpdateUserFields
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(userInput)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (us *UserServiceImpl) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user, err := checkIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

func checkIfExistId(id string) (models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func userExist(newUser models.CreateUserFields) *models.User {
	var existingUser models.User
	if err := database.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user := models.User{
				Name:      newUser.Name,
				Lastname:  newUser.Lastname,
				Username:  newUser.Username,
				Email:     newUser.Email,
				Password:  newUser.Password,
				CreatedAt: time.Now(),
			}
			return &user
		}
		return nil
	}

	return nil
}

func (newUser *CreateUserFields) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	newUser.Password = string(bytes)
	return nil
}

func (user *User) checkPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
