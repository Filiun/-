package router

import (
	"crm/api/login"
	"crm/api/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

	r := gin.Default()
	r.Use(Cors())
	gr:=r.Group("crm")
	testRouter(gr)
	r.Run(":8000")
}

func  testRouter(r *gin.RouterGroup)  {

	gr:=r.Group("/test")

	gr.GET("testedit", func(context *gin.Context) {
		context.String(http.StatusOK,"hello  world")
	})

	gr.POST("login",login.Index)
	gr.POST("add",article.Add)
	gr.GET("detail",article.Detail)
	gr.GET("list",article.List)
}

//跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
