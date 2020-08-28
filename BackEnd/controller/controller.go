package controller

import "github.com/gin-gonic/gin"

type Handler func(c *gin.Context) (int, interface{}, error)

func QueryExpr(c *gin.Context) (code int, data interface{}, err error) {
	data = map[string]interface{}{
		"sum": 123.44,
	}
	return
}

func Query(c *gin.Context) {
	code, data, err := QueryExpr(c)
	Res(c, code, data, err)
}
