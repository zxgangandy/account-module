package model

import (
	"account-module/pkg/baseerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

var R = NewResponse()

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    baseerr.Success.Code(),
		Message: baseerr.Success.Msg(),
		Data:    data,
	})
}

func (r *Response) Error(c *gin.Context, err error) {
	if err != nil {
		if v, ok := err.(*baseerr.Error); ok {
			response := Response{
				Code:    v.Code(),
				Message: v.Msg(),
				Data:    gin.H{},
			}

			c.JSON(v.StatusCode(), response)
			return
		}
	}

	c.JSON(http.StatusOK, Response{
		Code:    baseerr.Success.Code(),
		Message: baseerr.Success.Msg(),
		Data:    gin.H{},
	})
}
