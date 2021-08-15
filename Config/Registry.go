package Config

import (
	"Gogogo/Application/RouteHandler/Sample"
	"Gogogo/Application/RouteHandler/Welcome"
	"Gogogo/Packages/Handler"
)


var Registry = map[string] func() Handler.HandlerInterface{
	"AppWelcome": func() Handler.HandlerInterface {
		return Welcome.ApiHandler{}
	},
	"AppSample": func() Handler.HandlerInterface {
		return Sample.ApiHandler{}
	},
}
