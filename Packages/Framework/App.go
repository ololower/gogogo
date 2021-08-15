package App

import (
	"Gogogo/Config"
	"Gogogo/Packages/Request"
	"Gogogo/Packages/Route"
	"log"
	"net/http"
	"strings"
)

type App struct {
	routes []Route.Route

	request *http.Request
	responseWriter http.ResponseWriter
}

// InitApplication Основная инициализация приложения
// При инициализации выполняется загрузка роутов из всех валидных xml файлов в папке ./Routes
func (s App) InitRoutine() App {
	// Загружаем маршруты
	s.routes = Route.Loader{}.LoadRoutes()

	// Запуск основного обработчика для маршрутов приложения
	// Все маршруты проксируются через базовый /
	// и выбор метода для обработки маршрута будет через UrlMatcher
	http.HandleFunc("/", s.runHandleFunc)

	log.Fatal(http.ListenAndServe(":8080", nil))
	return s
}

// runHandleFunc Запускает функцию - обработчик для текущего роута и возвращает ответ
func (s App) runHandleFunc(w http.ResponseWriter, r *http.Request) {

	// Формируем appData для прокидования данных в обработчик маршрутов
	var appData = Request.NewAppData(r, w)

	//Выполняем поиск запрошенного url и вызываем соответствующий для этого метод
	for i := range s.routes {
		matcher := Route.NewUrlMatcher(
			s.routes[i].GetPathForHandler(),
			strings.Trim(appData.Request.URL.Path, "/"),
		)

		if matcher.Match() {
			var resp = Config.Registry[s.routes[i].Handler]().HandleMethod("GET", appData)
			resp.SetResponse(w)
		}
	}
}
