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
	shellPath := ctx.Param("name")
	shellPath = config.Config.HookPath + shellPath + ".sh"
	cmd := exec.Command(shellPath)
	args := []string{cmd.Args[0], config.Config.HookPath}
	cmd.Args = args
	err := cmd.Start()
	if err != nil {
		utils.Success(ctx, err)
		return
	}
	utils.Success(ctx, shellPath)
}
