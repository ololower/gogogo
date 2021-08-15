package Config

import (
	"Gogogo/Application/RouteHandler/Welcome"
)


var Registry = map[string] func() Welcome.ApiHandler{
	"AppWelcome": func() Welcome.ApiHandler {
		return Welcome.ApiHandler{}
	},
}
