package Route

import (
	"io/ioutil"
	"log"
)

// RouteLoader Загрузчик файлов с роутами
type Loader struct {

}

// LoadRoutes выполняет сканирование папки с роутами и передает загрузку конкретного роута
// соответсвующему классу
func (s Loader) LoadRoutes() []Route {
	var routes []Route

	var routesFolder = "./Routes"
	files, _ := ioutil.ReadDir(routesFolder)
	for i := range files {

		// Вызываем соответствующий лоадер из расширанея файла
		// Соглашение: структура, которая отвечает за загрузку роутов из файл должна
		// содежрать расширение файла
		// примеры:
		// расширение: xml, название структуры RouteXmlLoader
		// расширение: csv, название структуры RouteCsvLoader
		var filename = files[i].Name()
		loader, err := NewRouteLoaderForFile(files[i].Name())
		if err != nil {
			log.Fatal(err)
		}
		var filePath = routesFolder + "/" +filename
		loadedRoutes, _ := loader.LoadFileRoute(filePath)

		for i2 := range loadedRoutes {
			routes = append(routes, *loadedRoutes[i2])
		}
	}
	
	return routes
}
