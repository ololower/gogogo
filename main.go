package main

import App "Gogogo/Packages/Framework"


func main() {
	//http.HandleFunc("/", AppWelcome.ApiHandler{}.HandleGet)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	App.App{}.InitRoutine()

}
