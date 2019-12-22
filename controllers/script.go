/**
  create by yy on 2019-07-29
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"hook_scripts/libs"
	"hook_scripts/service"
)

func Hook(ctx *gin.Context) {

	shellName := ctx.Param("name")

	data := ctx.PostForm("payload")

	result := service.Hook(shellName, data)

	libs.Logger.Info(result)

	libs.Success(ctx, result)
}
