package AuthController

import (
	UserController "Users/controllers/User"
	"Users/httputil"
	"Users/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary Sign Up
// @Description add by json account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body models.SignUpData true "Sign Up account"
// @Success 200 {object} models.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /auth/signup [post]
func SignUp(c *gin.Context) {
	var input models.SignUpData
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	password, err := UserController.PasswordVerify(&input, c)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Name:     input.Name,
		Age:      input.Age,
		Email:    input.Email,
		Password: string(password),
	}

	if err := models.DB.Create(&user).Error; err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
