package controllers

import (
	"demoProject/internal/webapp/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func  HomeView(c *gin.Context)  {
	c.HTML(http.StatusOK,"/home/home.tmpl",nil)
}

func  LoginView(c *gin.Context)  {
	c.HTML(http.StatusOK,"/user/login.tmpl",nil)
}
func  RegirstView(c *gin.Context)  {
	c.HTML(http.StatusOK,"/user/regirst.tmpl",nil)
}

func UserLogin(c *gin.Context) {
	//message := c.PostForm("message")
	//nick := c.DefaultPostForm("nick", "anonymous")
	//c.HTML(http.StatusOK,"login.tmpl",gin.H{
	//	"message": "pong",
	//})
}
func UserRegirst(c *gin.Context) {
	var  userinfo models.Userinfo
	if err:=c.ShouldBindJSON(&userinfo);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"数据验证未通过",
		})
		return
	}
	flag, err := models.AddUser()
	if (err == nil&&flag==true){
		c.JSON(http.StatusOK,gin.H{
			"message":"添加成功" ,
		})
	}}



