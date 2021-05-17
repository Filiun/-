package article

import (
	"crm/db/modes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type  Param struct {
	Title  string `json:"title"  form:"title"  binding:"required"`
	Content  string `json:"test-editor-html-code"  form:"test-editor-html-code"  binding:"required"`
}

/**
添加博客文章
 */
func   Add(c *gin.Context)  {

	var  req Param
	if err:=c.ShouldBind(&req);err!= nil {
		c.JSON(0,gin.H{"msg":"确少参数","err":1})
		return
	}
	var  article  modes.Article
	article.Name=req.Title
	article.Content=req.Content
	article.Author="hhb"
	article.Uid=1
	article.Add()
	c.JSON(http.StatusOK,gin.H{"msg":"添加成功","err":0})
}

func  List(c *gin.Context)  {

	var  article  modes.Article
	list,err:=article.List(1)
	if err != nil {
		c.JSON(0,gin.H{"msg":"获取列表失败"+err.Error(),"err":1})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":"获取列表成功","err":0,"data":list})
}

func  Detail(c *gin.Context)  {

	id,_:=strconv.ParseInt(c.Query("id"),10,64)
	var  article  modes.Article
	err:=article.Info(id)

	fmt.Print(article)


	if err != nil  {
		c.JSON(0,gin.H{"msg":"获取详情失败"+err.Error(),"err":1})
		return
	}
	c.JSON(http.StatusOK,gin.H{"msg":"获取详情成功","err":0,"data":article})

}


