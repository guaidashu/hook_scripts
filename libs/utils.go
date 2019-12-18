/**
  create by yy on 2019/10/30
*/

package libs

import (
	"errors"
	"fmt"
	"hook_scripts/config"
	"runtime"
	"time"
)

func GetErrorString(err error) string {
	return fmt.Sprintf("error: %v", err)
}

func NewReportError(err error) error {
	if !config.Config.App.Debug {
		return err
	}
	_, fileName, line, _ := runtime.Caller(1)
	data := fmt.Sprintf("%v, report in: %v: in line %v", err, fileName, line)
	return errors.New(data)
}

func GetNowTimeMon(nowTimeStamp int64) string {
	if nowTimeStamp == 0 {
		nowTimeStamp = time.Now().Unix()
	}
	return time.Unix(nowTimeStamp, 0).UTC().Format("2006-01-02")
}

// 得到当前时间戳
func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}
