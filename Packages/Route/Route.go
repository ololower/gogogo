package Route

import (
	"Gogogo/Config"
	"net/http"
	"reflect"
	"strings"
)

type Route struct {
	Path	string
	Handler string
}

func NewRoute(path string, handler string) *Route {
	return &Route{Path: path, Handler: handler}
}

// Init инициализация маршрута в приложении
func (s Route) GetHandler(w http.ResponseWriter, r *http.Request) {

	v := Config.Registry[s.Handler]()

	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(w)
	inputs[1] = reflect.ValueOf(r)


	reflect.ValueOf(v).MethodByName("HandleGet").Call(inputs)

}

func (s Route) GetPathForHandler() string {
	var path = strings.Trim(s.Path, "/")
	return path
}