package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"package/Config"
	"package/Service"
)

var (
	cfgfile = "E:/package/config.yml"
)

func main() {
	Config.Initialize(cfgfile)
	service.Start()

	//写入gin日志并输出在控制台
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
