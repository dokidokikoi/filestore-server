package core

import (
	"encoding/json"
	"net/http"

	"file-store/errors"
)

type resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `josn:"data"`
}

func WriteResp(w http.ResponseWriter, data interface{}, err *errors.MyError) {
	var reply []byte
	tmp := resp{
		Code: err.Code,
		Msg:  err.Msg,
		Data: data,
	}
	reply, _ = json.Marshal(tmp)

	w.Write(reply)
}
