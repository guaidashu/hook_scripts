/**
  create by yy on 2019-07-29
*/

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hook_scripts/config"
	"hook_scripts/data_struct"
	"hook_scripts/libs"
	"os/exec"
	"strings"
)

func Hook(shellPath string, data string) string {

	var (
		ctx     context.Context
		cmd     *exec.Cmd
		err     error
		outPut  []byte
		payload data_struct.Payload
	)

	if err = json.Unmarshal([]byte(data), &payload); err != nil {
		return "get null commits content"
	}

	if config.Config.App.UpdateToken != "" &&
		payload.Commits[0].Message != "" &&
		!strings.Contains(payload.Commits[0].Message, config.Config.App.UpdateToken) {
		return "invalid update token"
	}

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
