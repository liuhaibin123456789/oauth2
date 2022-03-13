package tool

import (
	"aouth2Demo/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JsonOutPut(status int, err error, data interface{}, c *gin.Context) {
	if err != nil {
		errMod := model.Err{Status: status, ErrInfo: fmt.Sprintln(err)}
		c.JSON(200, errMod)
		return
	}
	c.JSON(200, model.Success{Data: data, Status: status})
}

//提交失误，补充注释作为修改，再次提交
