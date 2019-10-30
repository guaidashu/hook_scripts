/**
  create by yy on 2019-07-30
*/

package controllers

import (
	"fmt"
	"hook_scripts/libs"

	"github.com/gin-gonic/gin"
)

func TestParam(ctx *gin.Context) {
	fmt.Println(ctx.Param("name"))
	libs.Success(ctx, "ok")
}
