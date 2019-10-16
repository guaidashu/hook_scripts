/**
  create by yy on 2019-07-29
*/

package service

import (
	"context"
	"hook_scripts/config"
	"os/exec"
)

func Hook(shellPath string) string {
	var (
		ctx    context.Context
		cmd    *exec.Cmd
		err    error
		outPut []byte
	)
	ctx = context.TODO()
	shellPath = config.Config.Hook.Path + shellPath + config.Config.Hook.Suffix
	cmd = exec.CommandContext(ctx, shellPath)
	args := []string{cmd.Args[0], config.Config.Hook.Path}
	cmd.Args = args
	outPut, err = cmd.CombinedOutput()
	if err != nil {
		return "shell execute failed"
	}
	return string(outPut)
}
