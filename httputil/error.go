package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
	return
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
