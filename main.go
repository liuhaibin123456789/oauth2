package main

import (
	"aouth2Demo/api"
	_ "aouth2Demo/docs"
	middleware "aouth2Demo/middlewares"
	"aouth2Demo/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           aouth2实现
// @version         1.0
// @host            localhost:8084
// @description     aouth2学习
// @securityDefinitions.apikey CoreAPI
// @name Authorization
// @in header
func main() {

	router := gin.Default()
	err := tool.InitMysql()
	if err != nil {
		fmt.Println(fmt.Sprintf(err.Error()))
		return
	}
	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册swagger访问路由
	router.GET("/authorization-code", api.CreateCode)                         //获取授权码
	router.GET("/redirect", api.Code)                                         //该路由用于获取重定向后的code url参数
	router.POST("/access-token", api.PostToken)                               //通过授权码获取github的token
	//gin框架提供了对于原生http库的类型转化，这样可以方便地使用到标准库地中间件
	router.GET("/user-info",
		func(c *gin.Context) {
			token := c.Query("token")
			c.Request.Header.Add("token", token)
			fmt.Println(c.Request.Header.Values("token"))
			if token == "" {
				c.Abort()
			}
			c.Next()
		}, gin.WrapF(api.User)) //携带token获取用户信息,没有就创建
	//router.POST("user", api.CreateUser) //调用第三方资源获取的信息，实现注册
	err = router.Run(":8084")
	if err != nil {
		return
	}
}
