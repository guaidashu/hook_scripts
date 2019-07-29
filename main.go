/**
  create by yy on 2019-07-29
*/

package main

import (
	_ "hook_script/config"
	"hook_script/ginServer"
	_ "hook_script/router"
)

func main() {
	ginServer.Run("127.0.0.1:8081")
}
