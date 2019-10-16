/**
  create by yy on 2019-07-29
*/

package ginServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hook_scripts/config"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()
}

func GET(pattern string, function gin.HandlerFunc) {
	Router.GET(pattern, function)
}

func POST(pattern string, function gin.HandlerFunc) {
	Router.POST(pattern, function)
}

func Run(addr ...string) {
	if len(addr) > 0 || config.Config.App.Host == "" {
		_ = Router.Run(addr...)
		return
	}
	address := fmt.Sprintf("%v:%v", config.Config.App.Host, config.Config.App.Port)
	_ = Router.Run(address)
}
