//
package AppWelcomeWebHandler

import "fmt"

type WebHandler struct {}

// Обработка GET запроса в приложении
func (s WebHandler) HandleGetRoute()  {
	fmt.Print("props")
}



