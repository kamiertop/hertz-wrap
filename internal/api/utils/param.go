package utils

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"hertz/pkg/resp"
)

func ParamId(c *app.RequestContext) int {
	param := c.Param("id")
	if param == "" {
		resp.BadReqStr(c, "invalid id param")
		return 0
	}
	//for _, r := range param {
	//
	//}

	id, _ := strconv.Atoi(param)

	return id
}
