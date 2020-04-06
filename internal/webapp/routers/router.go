package routers

import (
	"demoProject/cmd/docs"
	"demoProject/internal/webapp/controllers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func SetRouter() *gin.Engine {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
	//router.Use(cors.InitCors)

	//api接口
	v1:=router.Group("api/v1")
	{
		user:=v1.Group("/user")
		{
			/*数据请求*/
			user.GET("",controllers.UserLogin)
			user.GET(":id",controllers.GetUserInfo)
			user.POST("/logindata",controllers.UserLogin)
			user.POST("/regirstdata",controllers.UserRegirst)
		}
	}

	user1:=router.Group("/user")
	{
		/*视图显示*/
		user1.GET("/login",controllers.LoginView)
	}
	router.GET("/home",controllers.HomeView)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
