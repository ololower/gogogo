package Welcome

import (
	"Gogogo/Packages/Request"
	"Gogogo/Packages/Response"
)

type ApiHandler struct {}

// Обработка GET запроса в приложении
func (s ApiHandler) HandleGet(appData *Request.AppData) Response.ResponseInterface {

	myJsonString := `{"some":"`+ appData.Request.Method +`"}`
	return Response.Json(myJsonString, 200)
}
