/**
  create by yy on 2019-07-29
*/

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hook_scripts/libs"
	"hook_scripts/service"
	"io/ioutil"
)

func Hook(ctx *gin.Context) {

	var (
		data []byte
		err  error
	)

	result := ctx.Param("name")

	if data, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
		err = libs.NewReportError(err)
		libs.Error(ctx, fmt.Sprintf("error: %v", err))
		return
	}

	libs.Success(ctx, service.Hook(result, data))
}
