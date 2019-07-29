/**
  create by yy on 2019-07-29
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"hook_script/utils"
)

func IvrHook(ctx *gin.Context) {
	utils.Success(ctx, "ok")
}
