package service

import (
	"ims/config"
	"ims/controller"
	"ims/model"
	"ims/orm"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func PreFlight(c *gin.Context) {
	origin := c.GetHeader("Origin")

	//c.Header("Content-Type", "Application/x-www-form-urlencoded;charset=UTF-8")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Cache-Control,Origin,Token,Content-Type")
		c.Header("Access-Control-Max-Age", "86400")
		c.Abort()

	} else {
		c.Header("Access-Control-Expose-Headers", "Token")
		c.Next()
	}
}

func Start() (err error) {
	if err = model.InitMySql(config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName); err != nil {
		orm.Close()
		return
	}

	router.Use(PreFlight)
	router.Use(PreFlight)

	router.OPTIONS("/login")
	router.POST("/login", controller.Login)

	// api := router.Group("/", controller.Auth)
	api := router.Group("/")
	{
		api.OPTIONS("/query")
		api.POST("/query", controller.Query)

		api.OPTIONS("/commit")
		api.POST("/commit")
	}

	if err = router.Run(config.ListenPort); err != nil {
		return
	}

	return
}
