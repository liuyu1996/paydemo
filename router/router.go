package router

import (
	"github.com/gin-gonic/gin"
	"paydemo/controller"
)

func Router() *gin.Engine{

	router := gin.Default()

	router.GET("/wxpay/test", controller.WxpayTestController)

	return router
}
