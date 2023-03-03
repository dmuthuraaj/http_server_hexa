package routers

import (
	"ex/internal/ports"
	"log"
	"net/http"
)

type RouterApp struct {
	handler ports.HandlerApp
}

func NewAdapter() *RouterApp {
	return &RouterApp{}
}

func middlewareLogging(logg http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request got from :%v Path:%s Method :%s", r.RemoteAddr, r.URL.Path, r.Method)
		logg.ServeHTTP(w, r)
	})
}

func (routera RouterApp) Handler() *http.ServeMux {

	router := http.NewServeMux()

	router.Handle("/", middlewareLogging(http.HandlerFunc(routera.handler.HomeHandler)))
	router.Handle("/signup", middlewareLogging(http.HandlerFunc(routera.handler.SignupHandler)))

	return router
}
