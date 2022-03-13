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
	router.POST("/authorization-code", api.PostCode)                          //获取授权码
	router.POST("/redirect", api.Code)                                        //该路由用于获取重定向后的code url参数
	/*
		router.POST("/access-token", PostToken)                                   //通过授权码获取github的token
		router.POST("/user-info", PostUser)                                       //携带token获取用户信息,没有就创建
	*/
	router.Run(":8084")
}

//提交失误，补充注释作为修改，再次提交
