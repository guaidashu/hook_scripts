/**
  create by yy on 2019-07-29
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"hook_scripts/service"
	"hook_scripts/utils"
)

func Hook(ctx *gin.Context) {
	result := ctx.Param("name")
	utils.Success(ctx, service.Hook(result))
}
