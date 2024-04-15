package api

import (
	"maps"
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

// routeTable storage handler route and it's name, this is not concurrent safe
// key: path
// value: routeName
var routeTable = make(map[string]string)

func SetRouteName(relativePath, routeName string) {
	routeTable[relativePath] = routeName
}

func GetRouteName(path string) (string, bool) {
	name, b := routeTable[path]

	return name, b
}

func GetRouteTable() map[string]string {
	return maps.Clone(routeTable)
}
