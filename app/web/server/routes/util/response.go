package util

import (
	"github.com/gin-gonic/gin"
	"github.com/heyujiang/shop/model/view"
	"github.com/heyujiang/shop/pkg/e"
	"net/http"
)

func SendResponse(c *gin.Context, code e.Code, data interface{}) {
	SendMsgResponse(c, code, code.ToDesc(), data)
}

func SendMsgResponse(c *gin.Context, code e.Code, msg string, data interface{}) {
	resp := view.Response{
		Success: code == e.Success,
		Code:    code,
		Msg:     msg,
		Data:    nil,
	}
	if resp.Success {
		resp.Data = data
	}
	c.JSON(http.StatusOK, resp)
}
