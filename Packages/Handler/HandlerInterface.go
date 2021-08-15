package Handler

import (
	"Gogogo/Packages/Request"
	"Gogogo/Packages/Response"
)

type HandlerInterface interface {
	HandleMethod(method string, appData *Request.AppData) Response.ResponseInterface
}
