package ports

import "net/http"

type RouterApp interface {
	Handler() *http.ServeMux
}

type HandlerApp interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
	HomeHandler(w http.ResponseWriter, r *http.Request)
}
