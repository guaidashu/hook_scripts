/**
  create by yy on 2019-07-29
*/

package service

import (
	"github.com/gin-gonic/gin"
	"hook_script/config"
	"hook_script/utils"
	"os/exec"
)

func Hook(ctx *gin.Context) {
	shellPath := ctx.Request.URL.Query().Get("name")
	shellPath = config.Config.HookPath + shellPath + ".sh"
	exec.Command(shellPath)
	utils.Success(ctx, shellPath)
}
