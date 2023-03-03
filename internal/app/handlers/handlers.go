package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type HandlerApp struct {
}

func NewAdapter() *HandlerApp {
	return &HandlerApp{}
}

func (handlerApp HandlerApp) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// var wg sync.WaitGroup

	if r.Method != "GET" {
		http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
		return
	}

	// wg.Add(1)
	// go func() {
	name := strings.Split(r.URL.Path, "/")
	err := json.NewEncoder(w).Encode("Hello " + name[1])
	if err != nil {
		http.Error(w, "unable to encode data", http.StatusInternalServerError)
		return
	}
	// 	defer wg.Done()
	// }()
	// wg.Wait()
}

func (handlerApp HandlerApp) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
		return
	}
}
