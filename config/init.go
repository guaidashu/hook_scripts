/**
  create by yy on 2019-07-29
*/

package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"hook_script/data_struct"
)

var Config data_struct.Config

func init() {
	fmt.Println("开始加载配置文件")
	err := configor.Load(&Config, "config/config.yml")
	if err != nil {
		fmt.Println("配置文件加载失败")
	} else {
		fmt.Println("配置文件加载完成")
	}
}
