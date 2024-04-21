package router

import "maps"

// routeTable storage handler route and it's name, this is not concurrent safe
// key: path
// value: routeName
var routeTable = make(map[[2]string]string)

func setRouteName(httpMethod, relativePath, routeName string) {
	routeTable[[2]string{httpMethod, relativePath}] = routeName
}

func GetRouteName(method, path string) (string, bool) {
	name, b := routeTable[[2]string{method, path}]

	return name, b
}

func GetRouteTable() map[[2]string]string {
	return maps.Clone(routeTable)
}
