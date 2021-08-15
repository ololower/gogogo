package Route

import "regexp"

type UrlMatcher struct {
	routePath 	string
	path 	string
}

func NewUrlMatcher(routePath string, path string) *UrlMatcher {
	return &UrlMatcher{routePath: routePath, path: path}
}

// Match проверяет, совпадает ли path с routePath
// routePath - путь, который настраивается в api.xml
// path - путь, который от открыл пользователь
func (s *UrlMatcher) Match() bool {
	matched, _ := regexp.MatchString(s.getRoutePathRegexp(), s.path)
	return matched
}

// GetMatches Возвращает меппинг динамических переменных из url
// todo сделать чтобы параметры можно было доставать по их названиям и удалить первый match
func (s UrlMatcher) GetMatches() []string {
	m1 := regexp.MustCompile(s.getRoutePathRegexp())
	res := m1.FindStringSubmatch(s.path)
	return res
}

func (s UrlMatcher) getRoutePathRegexp() string {
	var trailingVarRegexp = "\\{\\w+\\}"
	m1 := regexp.MustCompile(trailingVarRegexp)
	res := "^" + m1.ReplaceAllString(s.routePath, "(.*)") + "$"
	return res
	//return "^pages/(.*)/2020/(.*)$"
}
