package main

import (
	"firsthw/api"
	"net/http"
)

func main() {
	addr := ":8080"
	s := api.NewTask()
	api.CreateAndRunServer(s, addr)
	http.ListenAndServe(addr, nil)
}
