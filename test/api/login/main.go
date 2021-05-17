package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type  Param struct {
	Name  string `json:"name"  form:"name"  binding:"required"`
	Pass  string `json:"pass"  form:"pass"  binding:"required"`
}

func   Index(c *gin.Context)  {


	if err:=c.ShouldBind(&Param{});err!= nil {
		c.JSON(0,gin.H{"msg":"确少参数","err":1})
		return
	}
    c.JSON(http.StatusOK,gin.H{"msg":"登录成功","err":0})

}
