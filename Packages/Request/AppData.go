package Request

import "net/http"

type AppData struct {
	Request *http.Request
	responseWriter http.ResponseWriter
}

func NewAppData(request *http.Request, responseWriter http.ResponseWriter) *AppData {
	return &AppData{Request: request, responseWriter: responseWriter}
}
