package Route

import (
	"Gogogo/Packages/Route/Xml"
	"encoding/xml"
	"io/ioutil"
)

type XmlLoader struct {

}
// LoadFileRoute загрузка роутов для файлов типа xml
func (s XmlLoader) LoadFileRoute(filename string) ([]*Route, error)  {
	var routes []*Route

	byteVal, _ := ioutil.ReadFile(filename)

	var xmlRouteStructure Xml.RoutesListStructure

	xml.Unmarshal(byteVal, &xmlRouteStructure)

	for i := range xmlRouteStructure.XmlRoutes {
		routes = append(
			routes,
			NewRoute(xmlRouteStructure.XmlRoutes[i].Path, xmlRouteStructure.XmlRoutes[i].Handler),
		)
	}

	//fmt.Println(fmt.Sprintf("%#v", routes))
	//fmt.Print(routes)
	return routes, nil
}

