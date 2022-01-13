package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"package/Config"
	service "package/Service"
)

var (
	cfgfile = "E:/package_go/config.yml"
)

func init() {
	fmt.Println("我开始执行了!")
}

func main() {
	Config.Initialize(cfgfile)
	service.Start()

	//写入gin日志并输出在控制台
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//并发测试
	//go func() {
	//	i := 0
	//	for {
	//		i++
	//		fmt.Printf("new goroutine:i = d%\n", i)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main goroutine:i = d%\n", i)
	//	time.Sleep(time.Second)
	//	if i == 2 {
	//		break
	//	}
	//}
}
