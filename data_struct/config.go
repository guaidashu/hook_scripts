/**
  create by yy on 2019-07-29
*/

package data_struct

type Config struct {
	Hook HookConfig `json:"hook" xml:"hook"`
	App  AppConfig  `json:"app" xml:"app"`
}

type AppConfig struct {
	Host   string      `json:"host" xml:"host"`
	Port   interface{} `json:"port" xml:"port"`
	Debug  bool        `json:"debug"`
	LogDir string      `json:"log_dir" xml:"log_dir"`
}

type HookConfig struct {
	Path   string `json:"path" xml:"path"`
	Suffix string `json:"suffix" xml:"suffix"`
}
