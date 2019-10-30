/**
  create by yy on 2019/10/30
*/

package libs

import (
	"errors"
	"fmt"
	"hook_scripts/config"
	"runtime"
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
