package PackagesApp

import PackagesRoute "Gogogo/Packages/Route"

type App struct {

}

// InitApplication Основная инициализация приложения
func (s App) InitApplication() App {
	PackagesRoute.RouteLoader{}.LoadRoutes()
	return s
}