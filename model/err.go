package model

//Err 错误类型
type Err struct {
	Status  int    `json:"status"`   //状态码
	ErrInfo string `json:"err_info"` //错误信息
}

//提交失误，补充注释作为修改，再次提交
