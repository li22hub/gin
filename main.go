package main

import (
	"fmt"
	"package/Router"
)

type persion struct {
	name string
	city string
	age int
	address string
	tel int
}

func main() {
	//fmt.Println("hello world!")

	r := Router.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

	//var p1 persion
	//p1.name = "admin"
	//p1.city = "上海"
	//p1.age = 18
	//p1.address = "南京西路"
	//p1.tel = 15093029652
	//fmt.Printf("p1=%v\n",p1)

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
	//r := gin.Default()
	////限制上传最大尺寸
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/upload", func(c *gin.Context) {
	//	file, err := c.FormFile("file")
	//	if err != nil {
	//		c.String(500, "上传图片出错")
	//	}
	//	// c.JSON(200, gin.H{"message": file.Header.Context})
	//	c.SaveUploadedFile(file, file.Filename)
	//	c.String(http.StatusOK, file.Filename)
	//})
	//r.Run(":8889")

	//上传多个文件
	//r := gin.Default()
	//// 限制表单上传大小 8MB，默认为32MB
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/upload", func(c *gin.Context) {
	//	form, err := c.MultipartForm()
	//	if err != nil {
	//		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	//	}
	//	// 获取所有图片
	//	files := form.File["files"]
	//	// 遍历所有图片
	//	for _, file := range files {
	//		// 逐个存
	//		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	//			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	//			return
	//		}
	//	}
	//	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	//})
	////默认端口号是8080
	//r.Run(":8889")



}
