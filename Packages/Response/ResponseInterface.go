package Response

import "net/http"

type ResponseInterface interface {
	SetResponse(w http.ResponseWriter)
} 
