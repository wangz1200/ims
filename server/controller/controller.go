package controller

import "github.com/gin-gonic/gin"

type Handler func(c *gin.Context) (int, interface{}, error)
