/**
  create by yy on 2019-07-02
*/

package router

import (
	"fmt"
	"hook_script/controllers"
	"hook_script/ginServer"
)

func init() {
	fmt.Println("开始初始化router")
	ginServer.GET("/hook", controllers.Hook)
	fmt.Println("router初始化成功")
}
