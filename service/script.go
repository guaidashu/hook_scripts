/**
  create by yy on 2019-07-29
*/

package service

import (
	"hook_scripts/config"
	"os/exec"
)

func Hook(shellPath string) string {
	shellPath = config.Config.Hook.Path + shellPath + config.Config.Hook.Suffix
	cmd := exec.Command(shellPath)
	args := []string{cmd.Args[0], config.Config.Hook.Path}
	cmd.Args = args
	err := cmd.Start()
	if err != nil {
		return "shell execute failed"
	}
	return shellPath
}
