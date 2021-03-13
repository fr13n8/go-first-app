package UserController

import (
	"Users/models"
	"errors"
	"net/http"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
	return
}

func GetUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
	return
}

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

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdatedData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(&models.User{Age: input.Age})
	c.JSON(http.StatusOK, gin.H{"data": "Successfuly updated!"})
	return
}

func PasswordVerify(input *models.SignUpData, c *gin.Context) ([]byte, error) {
	if input.Password != input.PasswordConfirm {
		return nil, errors.New("Password doesnt match")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return password, nil
}

func JwtGenerate(u models.User) ([]byte, error) {
	var hs = jwt.NewHS256([]byte("secret"))
	now := time.Now()
	pl := models.JwtPayload{
		Payload: jwt.Payload{
			Issuer:         "gbrlsnchs",
			Subject:        "someone",
			Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
		Name:  u.Name,
		Age:   int(u.Age),
		Email: u.Email,
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return token, nil
}

func JwtVerify(token []byte) (jwt.Header, error) {
	var hs = jwt.NewHS256([]byte("secret"))
	var pl models.JwtPayload
	hd, err := jwt.Verify(token, hs, &pl)
	if err != nil {
		return jwt.Header{}, errors.New(err.Error())
	}

	return hd, nil
}
