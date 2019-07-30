/**
  create by yy on 2019-07-30
*/

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hook_script/utils"
)

func TestParam(ctx *gin.Context) {
	fmt.Println(ctx.Param("name"))
	utils.Success(ctx, "ok")
}
