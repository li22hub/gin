package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//fmt.Println("hello world!")

	//基本路由
	//r := gin.Default()
	//r.GET("/admin", func(c *gin.Context) {
	//	c.String(http.StatusOK,"hello world")
	//})
	//r.Run(":8888")

	//获取get api参数
	//r := gin.Default()
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	action = strings.Trim(action,"/")
	//	c.String(http.StatusOK,name+" is "+action)
	//})
	//r.Run(":8889")

	//获取url参数
	//r := gin.Default()
	//r.GET("/user", func(c *gin.Context) {
	//	name := c.DefaultQuery("name","迪迦")
	//	c.String(http.StatusOK,fmt.Sprintf("hello %s", name))
	//})
	//r.Run(":8889")

	//表单参数
	//r := gin.Default()
	//r.POST("/form", func(c *gin.Context) {
	//	types := c.DefaultPostForm("type","post")
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	c.String(http.StatusOK,fmt.Sprintf("username:%s,password:%s,type:%s",username,password,types))
	//})
	//r.Run(":8889")

	//上传单个文件
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()

}
