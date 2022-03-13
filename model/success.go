package model

// Success 服务成功返回类型
type Success struct {
	Status int         `json:"status"` //状态码
	Data   interface{} `json:"data"`   //数据源
}

//AcceptGitHubToken 申请github token是的数据类型
type AcceptGitHubToken struct {
	AccessToken string ` json:"access_token"` //token字符串
	Scope       string `json:"scope"`         //
	TokenType   string `json:"token_type"`    //token类型,一般放在Authorization的bearer字符串后面
}

//AcceptGitHubUser 用户信息数据类型
type AcceptGitHubUser map[string]interface{}
