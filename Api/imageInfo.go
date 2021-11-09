package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

//测试gin
func Test(ctx *gin.Context) {
	fmt.Println("ok")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "21世纪的c语言",
	})
}

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//上传file图片
func UpFile(ctx *gin.Context) {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 //限制上传最大尺寸
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(500, "图片file丢失")
	}
	t := time.Now().Format("20060102")
	_dir := "./Image/" + t
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		fmt.Printf("has dir![%v]\n", _dir)
	} else {
		fmt.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	ctx.SaveUploadedFile(file, _dir+"/"+file.Filename)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "上传成功了 大兄弟",
		"path":    _dir + "/" + file.Filename,
	})
}
