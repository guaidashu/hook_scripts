/**
  create by yy on 2019-08-23
*/

package libs

import (
	"fmt"
	"github.com/op/go-logging"
	"hook_scripts/config"
	"io"
	"os"
)

var Logger *logging.Logger

func InitLogger() (error, string) {
	ok, _ := PathExists(config.Config.App.LogDir)
	if !ok {
		err := os.MkdirAll(config.Config.App.LogDir, 0777)
		if err != nil {
			return err, fmt.Sprintf("mkdir failed, error: %v", err)
		}
	}
	LogFormat := "%{color}%{level:.4s}:%{time:2006-01-02 15:04:05.000}[%{id:03x}] %{shortfile} %{color:reset} %{message}"
	Logger = logging.MustGetLogger("")
	logFile := GetNowTimeMon(GetNowTimeStamp()) + ".log"
	writeFd, err := os.OpenFile(config.Config.App.LogDir+"/"+logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println()
		return err, fmt.Sprintf("open file[%s] failed[%s]", logFile, err)
	}
	format := logging.MustStringFormatter(LogFormat)
	backendList := make([]logging.Backend, 0)
	for _, out := range []io.Writer{writeFd, os.Stdout} {
		backendInfo := logging.NewLogBackend(out, "", 0)
		backendInfoFormatter := logging.NewBackendFormatter(backendInfo, format)
		backendInfoLeveld := logging.AddModuleLevel(backendInfoFormatter)
		backendInfoLeveld.SetLevel(logging.INFO, "")

		// switch "INFO" {
		// case "ERROR":
		//	backendInfoLeveld.SetLevel(logging.ERROR, "")
		// case "WARNING":
		//	backendInfoLeveld.SetLevel(logging.WARNING, "")
		// case "NOTICE":
		//	backendInfoLeveld.SetLevel(logging.NOTICE, "")
		// case "INFO":
		//	backendInfoLeveld.SetLevel(logging.INFO, "")
		// case "DEBUG":
		//	backendInfoLeveld.SetLevel(logging.DEBUG, "")
		// default:
		//	backendInfoLeveld.SetLevel(logging.ERROR, "")
		// }

		bk := backendInfoLeveld
		backendList = append(backendList, bk)
	}

	logging.SetBackend(backendList...)

	return nil, ""
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
