package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"package/Database"
	"time"
)

//写入队列
func PingMQsend(ctx *gin.Context) {
	// amqp://用户名:密码@地址:端口号/host
	err := Database.SetupRMQ("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println("mq初始化出错 : ", err.Error())
	}
	err = Database.Ping()
	if err != nil {
		fmt.Println("mq测试链接(ping)出错 : ", err.Error())
	}

	//入队操作
	for i := 0; i < 10; i++ {
		err = Database.Publish("first", "当前时间："+time.Now().String())
		if err != nil {
			fmt.Println("入队出错 : ", err.Error())
		}
		//time.Sleep(1 * time.Second)
	}
	Database.Close()
	ctx.JSON(http.StatusOK,"写入队列成功")
}

//消费
func PingMQreceive(ctx *gin.Context)  {
	err := Database.SetupRMQ("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println("mq初始化出错 : ", err.Error())
	}
	err = Database.Ping()
	if err != nil {
		fmt.Println("mq测试链接(ping)出错 : ", err.Error())
	}

	fmt.Println("receive message")
	err = Database.Receive("first", "second", func(msg *string) {
		fmt.Printf("receve msg is :%s\n", *msg)
	})
	if err != nil {
		fmt.Println("消费出错 : ", err.Error())
	}
	ctx.JSON(http.StatusOK,"已消费完毕")
}