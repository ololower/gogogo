package Response

import "net/http"

type JsonResponse struct {
	content string
	responseCode int
}

func NewJsonResponse(content string, responseCode int) ResponseInterface {
	return &JsonResponse{content: content, responseCode: responseCode}
}

func (s JsonResponse) SetResponse(w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(s.content))
}
