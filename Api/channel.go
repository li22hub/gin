package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把sum发送到通道c
}

func ChannelT(ctx *gin.Context) {
	s := []int{7, 2, 8, -6, 1, 0}
	c := make(chan int)
	go sum(s, c)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	C, B, A := <-c, <-c, <-c //从通道c中接受
	ctx.JSON(http.StatusOK, gin.H{
		"A": A, //-5
		"B": B, //17
		"C": C, //12
	})
}
