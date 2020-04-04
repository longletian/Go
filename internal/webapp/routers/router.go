package routers

import (
	"demoProject/internal/webapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	//加载静态文件
	router.Static("/static","../internal/webapp/assets")
	//加载模板
	router.LoadHTMLGlob("../internal/webapp/template/*")
	/*
	router.SetFuncMap(template.FuncMap{
		 "safe": func(str string) template.HTML{
			 return template.HTML(str)
		 },
	})
	router.LoadHTMLFiles("../internal/webapp/template/login.tmpl")

	 */

	//用户组
	user:=router.Group("/user")
	{
		/*视图显示*/
		user.GET("/login",controllers.LoginView)
		user.GET("/regirst",controllers.RegirstView)

		/*数据请求*/
		user.POST("/logindata",controllers.UserLogin)
		user.POST("/regirstdata",controllers.UserRegirst)
	}

	router.GET("/home",controllers.HomeView)
	return router
}