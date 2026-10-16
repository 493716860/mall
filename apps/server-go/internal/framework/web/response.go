package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommonResult 1:1 对齐 芋道 CommonResult<T>
type CommonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// PageResult 1:1 对齐 芋道 PageResult<T>
type PageResult struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, CommonResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// PageSuccess 分页成功响应
func PageSuccess(c *gin.Context, list any, total int64) {
	if list == nil {
		list = []any{}
	}
	c.JSON(http.StatusOK, CommonResult{
		Code: 0,
		Msg:  "success",
		Data: PageResult{
			List:  list,
			Total: total,
		},
	})
}

// Error 失败响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, CommonResult{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// Unauthorized 401 响应
func Unauthorized(c *gin.Context, msg string) {
	if msg == "" {
		msg = "账号未登录或 Token 已过期"
	}
	c.JSON(http.StatusOK, CommonResult{
		Code: 401,
		Msg:  msg,
		Data: nil,
	})
}
