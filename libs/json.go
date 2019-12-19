/**
  create by yy on 2019/12/19
*/

package libs

import (
	"bytes"
	"encoding/json"
)

func Marshal(v interface{}, isEscape ...bool) (data []byte, err error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	if len(isEscape) > 0 {
		jsonEncoder.SetEscapeHTML(isEscape[0])
	}
	err = jsonEncoder.Encode(v)
	data = bf.Bytes()
	return
}
