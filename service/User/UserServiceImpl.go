package services

import (
	"net/http"
	"time"
	database "user_auth/database/config"
	models "user_auth/models/User"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	c       *gin.Context
	newUser models.CreateUserFields
	user    models.User
}

func NewUserServiceImpl(c *gin.Context, newUser models.CreateUserFields, user models.User) UserService {
	return &UserServiceImpl{c, newUser, user}
}

func (service *UserServiceImpl) GetSpecificUser(c *gin.Context) {
	id := c.Param("id")
	user, err := checkIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (service *UserServiceImpl) RegisterUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&service.newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := service.hashPassword(service.newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	user := userExist(service.newUser)
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

func (service *UserServiceImpl) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := checkIfExistId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (service *UserServiceImpl) DeleteUser(c *gin.Context) {
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
	var user models.User
	if err := database.DB.Where("email = ?", newUser.Email).First(&user).Error; err != nil {
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

func (us *UserServiceImpl) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	us.newUser.Password = string(bytes)
	return nil
}

func (us *UserServiceImpl) checkPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(us.user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
