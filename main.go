/**
  create by yy on 2019-07-29
*/

package main

import (
	_ "hook_scripts/config"
	"hook_scripts/ginServer"
	_ "hook_scripts/router"
)

func main() {
	ginServer.Run("127.0.0.1:8099")
}
