package MessageController

import (
	"Users/httputil"
	"Users/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendMessage godoc
// @Summary Send Message
// @Description send message
// @Tags message
// @Accept  json
// @Produce  json
// @Param message body models.NewMessage true "Send message"
// @Success 200 {object} models.Message
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /message/send [post]
func SendMessage(c *gin.Context) {
	var message *models.NewMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := models.DB.Create(&message).Error; err != nil {
		log.Printf("Error: %v", err)
		return
	}

	c.JSON(http.StatusOK, message)
}
