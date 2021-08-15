package Response

// Json Возвращает ответ в формате json
func Json(content string, responseCode int) ResponseInterface  {
	return NewJsonResponse(content, responseCode)
}
