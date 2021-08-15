package Response

func Json(content string, responseCode int) ResponseInterface  {
	return NewJsonResponse(content, responseCode)
}
