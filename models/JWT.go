package models

import (
	"errors"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

func (u *User) JwtGenerate(c *gin.Context) ([]byte, error) {
	var hs = jwt.NewHS256([]byte("secret"))
	now := time.Now()

	pl := &JwtPayload{
		Payload: jwt.Payload{
			Issuer:         u.Email,
			Subject:        c.Request.Header.Get("Origin"),
			Audience:       jwt.Audience{c.Request.Header.Get("Origin")},
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          u.Name,
		},
	}

	token, err := jwt.Sign(pl, hs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return token, nil
}
