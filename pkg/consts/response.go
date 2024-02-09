package consts

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

const (
	_errParam        = "invalid params"
	_errUnauthorized = "unauthorized"
)

type Response struct {
	Msg string `json:"msg,omitempty"`
	Err string `json:"err,omitempty"`
}

func InterServerError(c *app.RequestContext, msg string, err error) {
	c.JSON(consts.StatusInternalServerError, Response{
		Msg: msg,
		Err: err.Error(),
	})
	// 中间件去捕获Error
	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: msg,
	})
}

func BadRequest(c *app.RequestContext, err error) {
	c.JSON(consts.StatusBadRequest, Response{
		Msg: _errParam,
		Err: err.Error(),
	})

	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: _errParam,
	})
}

func Unauthorized(c *app.RequestContext, err error) {
	c.JSON(consts.StatusUnauthorized, Response{
		Msg: _errUnauthorized,
		Err: err.Error(),
	})

	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: _errParam,
	})
}
