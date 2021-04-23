package AuthController

import (
	UserController "Users/controllers/User"
	"Users/httputil"
	"Users/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
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
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			httputil.NewError(c, http.StatusBadRequest, fieldErr)
			return // exit on first error
		}
	}

	password, err := UserController.PasswordVerify(&input)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	user := &models.User{
		Name:     input.Name,
		Age:      input.Age,
		Email:    input.Email,
		Password: string(password),
		Verified: false,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		var mysqlErr = err.(*mysql.MySQLError)
		if mysqlErr.Number == 1062 {
			httputil.NewError(c, http.StatusInternalServerError, errors.New("Key: 'SignUpData.Email' Error:This email already exists."))
		}
		return
	}

	go func() {
		if err := user.SendVerifyEmail(c); err != nil {
			log.Println(err.Error())
		}
	}()

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Verify(c *gin.Context) {
	pl, err := UserController.JwtVerify([]byte(c.Param("token")))
	if err != nil {
		log.Printf("Error verifying token: %v", err)
		httputil.NewError(c, http.StatusNotFound, errors.New("Bad token"))
		return
	}

	var user *models.User
	email := pl.Issuer

	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("User not found:%v", err)
		httputil.NewError(c, http.StatusNotFound, errors.New("User not found"))
		return
	}

	models.DB.Model(&user).Updates(&models.User{Verified: true})
	c.Redirect(http.StatusFound, pl.Subject)
}
