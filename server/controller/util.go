package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, code int, data interface{}, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}

	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Auth(c *gin.Context) {
	tokenStr := c.GetHeader("Token")

	if token, err := parseToken(tokenStr); err != nil {
		Res(c, -1, nil, errors.New("Token令牌错误！"))
		c.Abort()

	} else {
		c.Header("Token", tokenStr)
		c.Set("Token", token)
		c.Next()
	}
}
