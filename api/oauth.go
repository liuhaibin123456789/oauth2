package api

import (
	"aouth2Demo/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// CreateCode
// @summary  获取github授权码（没有找到swaggo对应重定向得注解耶。swagger调试接口还不行，所以只能使用postman或者浏览器测试了）
// @produce  json
// @Param    client_id     query  string  true  "开放平台申请的应用ID"
// @param    redirect_uri  query  string  false  "授权成功后，跳转的地址,GitHub已配置回调url，此处为可选"
// @router   /authorization-code [get]
func CreateCode(c *gin.Context) {
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
// @router   /redirect [get]
func Code(c *gin.Context) {
	code := c.Query("code")
	tool.JsonOutPut(200, nil, gin.H{"code": code}, c)
}

// PostToken
// @summary  通过code拿取github颁发的token
// @produce  json
// @Param    code     query  string  true  "授权码"
// @Param    client_id     query  string  true  "授权客户端id"
// @Param    client_secret     query  string  true  "客户端密钥"
// @Param    redirect_uri     query  string  false  "用户获得授权后被发送到的应用程序中的 URL"
// @success  200       {object}  model.AcceptGitHubToken  "成功"
// @failure  200       {object}  model.Err  "请求错误"
// @router   /access-token [post]
func PostToken(c *gin.Context) {
	//必选
	code := c.Query("code")
	clientId := c.Query("client_id")
	clientSecret := c.Query("client_secret")
	//可选
	redirectUri := c.Query("redirect_uri")
	fmt.Println(c.Request.Method)
	c.Redirect(http.StatusMovedPermanently, "https://github.com/login/oauth/access_token?client_id="+clientId+"&client_secret="+clientSecret+"&code="+code+"&redirect_uri="+redirectUri)
}

//// User
//// @summary  通过token拿取资源
//// @produce  json
//// @success  200       {object}  model.AcceptGitHubUser  "成功"
//// @failure  200       {object}  model.Err  "请求错误"
//// @Security CoreAPI
//// @router   /user-info [get]
//func User(c *gin.Context) {
//	//初步查看并校验token
//	authorization := c.Request.Header.Get("Authorization")
//	if authorization == "" {
//		tool.JsonOutPut(200, errors.New("the Authorization is empty"), "", c)
//		return
//	}
//	parts := strings.SplitN(authorization, " ", 2)
//	if !(len(parts) == 2 && parts[0] == "Bearer") {
//		tool.JsonOutPut(200, errors.New("authorization header is wrong,no bearer"), "", c)
//		return
//	}
//	//转化header bearer为token
//	//c.Header("Authorization", "Token "+parts[1])
//	//重定向
//	c.Redirect(308, "https://api.github.com/user")
//}

func User(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var err error
	//请求
	r1, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		fmt.Println("request is wrong")
		return
	}
	////通过url 传token
	//token := r.URL.Query().Get("token")
	//fmt.Println(r.URL, r.URL.Query(), token)
	//设置请求头
	r1.Header.Set("accept", "application/json")
	token := r.Header.Get("token")
	r1.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	//模拟客户端发起请求
	client := http.Client{}
	var res *http.Response
	if res, err = client.Do(r1); err != nil {
		fmt.Println(fmt.Sprintf(err.Error()))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	//取数据
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		fmt.Println(fmt.Sprintf(err.Error()))
		return
	}
	//返回json
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//func CreateUser(c *gin.Context) {
//	//保存用户信息至数据库
//
//}
