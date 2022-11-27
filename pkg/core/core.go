package core

import (
	"file-store/pkg/errors"

	"github.com/gin-gonic/gin"
)

type resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `josn:"data"`
}

func WriteResp(c *gin.Context, data interface{}, err *errors.MyError) {
	c.JSON(err.Code, resp{
		Msg:  err.Msg,
		Data: data,
	})
}
