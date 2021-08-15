package Route

type UrlMatcher struct {
	routePath 	string
	path 	string
}

func NewUrlMatcher(routePath string, path string) *UrlMatcher {
	return &UrlMatcher{routePath: routePath, path: path}
}

// match проверяет, совпадает ли path с routePath
func (s UrlMatcher) Match() bool {
	return s.routePath == s.path
}
