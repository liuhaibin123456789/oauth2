package api

import (
	"aouth2Demo/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostCode
// @summary  获取github授权码
// @produce  json
// @Param    client_id     query  string  true  "开放平台申请的应用ID"
// @param    redirect_uri  query  string  true  "授权成功后，跳转的地址"
// @success  200       {object}  model.Success  "成功"
// @failure  200       {object}  model.Err  "请求错误"
// @router   /authorization-code [post]
func PostCode(c *gin.Context) {
	//使用url参数（client_id，redirect_uri访问GitHub授权码api）
	clientID := c.Query("client_id")
	redirectUri := c.Query("redirect_uri")
	//重定向
	c.Redirect(http.StatusMovedPermanently, "https://github.com/login/oauth/authorize?client_id="+clientID+"&redirect_uri="+redirectUri)
}

// Code
// @summary  获取重定向后的code url参数
// @produce  json
// @Param    code     query  string  true  "授权码"
// @success  200       {object}  model.Success  "成功"
// @failure  200       {object}  model.Err  "请求错误"
// @router   /redirect [post]
func Code(c *gin.Context) {
	code := c.Query("code")
	tool.JsonOutPut(200, nil, gin.H{"code": code}, c)
}

//提交失误，补充注释作为修改，再次提交
