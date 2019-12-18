/**
  create by yy on 2019-07-29
*/

package service

import (
	"context"
	"fmt"
	"hook_scripts/config"
	"hook_scripts/libs"
	"os/exec"
)

func Hook(shellPath string, data []byte) string {
	var (
		ctx    context.Context
		cmd    *exec.Cmd
		err    error
		outPut []byte
	)

	return string(data)

	ctx = context.TODO()

	shellPath = config.Config.Hook.Path + shellPath + config.Config.Hook.Suffix

	cmd = exec.CommandContext(ctx, shellPath)

	args := []string{cmd.Args[0], config.Config.Hook.Path}

	cmd.Args = args

	outPut, err = cmd.CombinedOutput()

	if err != nil {
		return fmt.Sprintf("shell execute failed: %v", libs.NewReportError(err).Error())
	}

	return string(outPut)
}
