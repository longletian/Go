package main

import (
"demoProject/internal/webapp/routers"
"github.com/gin-gonic/gin"
"io"
"os"
)

func main() {
	//终端主函数
	router:=routers.SetRouter()
	f,_:=os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f,os.Stdout)
	router.Run("8088")
}
