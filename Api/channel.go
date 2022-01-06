package Api

import (
	"fmt"
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

//通道
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

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//关闭通道
func ChannelClose(ctx *gin.Context) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
