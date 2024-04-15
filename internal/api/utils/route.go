package utils

import (
	"maps"
)

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
