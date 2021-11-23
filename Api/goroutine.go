package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Hello(ctx *gin.Context) {
	go say("hello world!!!")
	say("hello")
}
