/**
  create by yy on 2019/12/21
*/

package data_struct

type Payload struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
}
