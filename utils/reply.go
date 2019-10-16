/**
  create by yy on 2019-07-29
*/

package utils

import (
	"github.com/gin-gonic/gin"
	"hook_scripts/data_struct"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	r := &data_struct.Reply{
		Code: 0,
		Msg:  "",
		Data: data,
	}
	ctx.JSON(http.StatusOK, r)
}

func Error(ctx *gin.Context, msg string) {
	r := &data_struct.Reply{
		Code: 1,
		Msg:  msg,
		Data: "",
	}
	ctx.JSON(http.StatusOK, r)
}
