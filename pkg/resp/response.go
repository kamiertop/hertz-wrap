package resp

import (
	stderr "errors"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

const (
	_errParam        = "invalid params"
	_errUnauthorized = "unauthorized"
	_msg             = "msg"
	_ok              = "ok"
	_data            = "data"
	_err             = "err"
)

type resp = map[string]any

func InterServerError(c *app.RequestContext, msg string, err error) {
	c.JSON(consts.StatusInternalServerError, resp{
		_msg: msg,
		_err: err.Error(),
	})
	// 中间件去捕获Error
	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: msg,
	})
}

func BadRequest(c *app.RequestContext, err error) {
	c.JSON(consts.StatusBadRequest, resp{
		_msg: _errParam,
		_err: err.Error(),
	})

	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: _errParam,
	})
}

func BadReqStr(c *app.RequestContext, msg string) {
	c.JSON(consts.StatusBadRequest, resp{
		_msg: _errParam,
		_err: stderr.New(msg),
	})

	c.Errors = append(c.Errors, &errors.Error{
		Err:  stderr.New(msg),
		Type: errors.ErrorTypeAny,
		Meta: _errParam,
	})
}

func Unauthorized(c *app.RequestContext, err error) {
	c.JSON(consts.StatusUnauthorized, resp{
		_msg: _errUnauthorized,
		_err: err.Error(),
	})

	c.Errors = append(c.Errors, &errors.Error{
		Err:  err,
		Type: errors.ErrorTypeAny,
		Meta: _errParam,
	})
}

func SuccessOK(c *app.RequestContext) {
	c.JSON(http.StatusOK, resp{
		_msg: _ok,
	})
}

func SuccessData(c *app.RequestContext, data any) {
	c.JSON(http.StatusOK, resp{
		_msg:  _ok,
		_data: data,
	})
}
