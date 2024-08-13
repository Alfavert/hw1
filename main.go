package main

import (
	"firsthw/api"
	_ "firsthw/docs"
	"net/http"
)

// @title My RESTfull api
// @version 1.0
// @description This is a sample server of testing Swagger.

// @host localhost:8000
// @BasePath /
func main() {
	addr := ":8000"
	s := api.NewTask()
	api.CreateAndRunServer(s, addr)
	http.ListenAndServe(addr, nil)
}
