package web

import (
	"modapilab1/internal/controller"
	"net/http"
)

type Router struct {
	FindDataHandler controller.FindDataController
}

func (r *Router) StartRouter() {
	http.HandleFunc("/", r.FindDataHandler.FindData)
	http.ListenAndServe(":8080", nil)
}
