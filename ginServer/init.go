/**
  create by yy on 2019-07-29
*/

package ginServer

import "github.com/gin-gonic/gin"

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
	_ = Router.Run(addr...)
}
