package AuthController

import (
	"Users/httputil"
	"Users/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SignIn godoc
// @Summary Sign In
// @Description add by json account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body models.SignInData true "Sign In account"
// @Success 200 {object} models.SignInResponseData
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /auth/signin [post]
func SignIn(c *gin.Context) {
	var input models.SignInData
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("User not found."))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Wrong password."))
		return
	}

	if !user.Verified {
		httputil.NewError(c, http.StatusForbidden, errors.New("Please verify your account"))
		return
	}

	token, err := user.JwtGenerate(c)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, &models.SignInResponseData{
		User: user, Token: string(token),
	})
}
