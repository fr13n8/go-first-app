package UserController

import (
	"Users/httputil"
	"Users/models"
	"errors"
	"net/http"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ListUsers godoc
// @Summary List users
// @Tags users
// @Description Get users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
	return
}

// GetUser godoc
// @Summary Get user details
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	return
}

// DeleteUser godoc
// @Summary Delete an user
// @Description Delete by user ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID" Format(int64)
// @Success 204 {string} string "status"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": "Deleted successfuly!"})
	return
}

// UpdateAccount godoc
// @Summary Update an user
// @Description Update by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Param  user body models.UpdatedData true "Update user"
// @Success 200 {object} models.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user *models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdatedData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(&models.User{Age: input.Age, Name: input.Name, Email: input.Email})
	c.JSON(http.StatusOK, gin.H{"data": &user})
	return
}

func PasswordVerify(input *models.SignUpData) ([]byte, error) {
	if input.Password != input.PasswordConfirm {
		return nil, errors.New("Password doesnt match")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return password, nil
}

func JwtVerify(token []byte) (models.JwtPayload, error) {
	var hs = jwt.NewHS256([]byte("secret"))
	var pl models.JwtPayload
	_, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		return models.JwtPayload{}, errors.New(err.Error())
	}

	return pl, nil
}
