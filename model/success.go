package model

type Success struct {
	Status int         `json:"status"` //状态码
	Data   interface{} `json:"data"`   //数据源
}

//提交失误，补充注释作为修改，再次提交
