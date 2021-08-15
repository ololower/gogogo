package Route

import (
	"errors"
	"path/filepath"
	"strings"
)

// LoaderInterface интерфейс для загрузки роутов из файлов
type LoaderInterface interface {
	LoadFileRoute(filename string) ([]*Route, error)
}

// NewRouteLoader возвращает инстанс нового загрузчика роутов
func NewRouteLoaderForFile(filename string) (LoaderInterface, error) {

	var fileExtension = strings.Trim(filepath.Ext(filename), ".")

	var extensionInLowercase = strings.ToLower(fileExtension)

	if extensionInLowercase == "xml" {
		return XmlLoader{}, nil
	}

	var error = errors.New("Route loader for filetype " + extensionInLowercase + " is not exits")

	return nil, error
}

