package Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "package/Common"
	"package/Models"
)

func FindAll(ctx *gin.Context) {
	var resData common.ResponseList
	data := []*Models.TestDetail{}
	data, count, err := Models.List(data)
	if err != nil {
		resData.Code = 500
		resData.Msg = "失败"
		ctx.JSON(http.StatusInternalServerError, resData)
		return
	}
	resData.Code = 200
	resData.Msg = "成功"
	resData.Data = data
	resData.TotalCount = count
	ctx.JSON(http.StatusOK, resData)
	return
}
