package Welcome

import (
	"Gogogo/Packages/Request"
	"Gogogo/Packages/Response"
	"strings"
)

type ApiHandler struct {}

// Обработка GET запроса в приложении
func (s ApiHandler) HandleGet(appData *Request.AppData) Response.ResponseInterface {
	myJsonString := `{"some":"`+ appData.Request.Method +`"}`
	return Response.Json(myJsonString, 200)
}

// HandleMethod определяет метод, который будет обрабатывать различные типы запросов
// todo настроить методы по умолчанию
// get - HandleGet
// post - HandlePost
// put - HandlePut
// delete - HandleDelete
func (s ApiHandler) HandleMethod(method string, appData *Request.AppData) Response.ResponseInterface {
	 var handlers = map[string] func() Response.ResponseInterface {
		"get": func() Response.ResponseInterface {
			return s.HandleGet(appData)
		},
	}

	var methodInLovercase = strings.ToLower(method)
	return handlers[methodInLovercase]()
}