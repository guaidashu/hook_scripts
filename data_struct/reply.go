/**
  create by yy on 2019-07-29
*/

package data_struct

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
