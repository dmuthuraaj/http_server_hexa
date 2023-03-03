package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type HandlerApp struct {
}

func NewAdapter() *HandlerApp {
	return &HandlerApp{}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	if r.Method != "GET" {
		http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
		return
	}

	wg.Add(1)
	go func() {
		log.Printf("Request Method :%v", r.Method)
		log.Printf("Request Host :%v", r.Host)
		log.Printf("Remote Address :%v", r.RemoteAddr)

		name := strings.Split(r.URL.Path, "/")
		err := json.NewEncoder(w).Encode("Hello " + name[1])
		if err != nil {
			http.Error(w, "unable to encode data", http.StatusInternalServerError)
			return
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not implemented", http.StatusMethodNotAllowed)
		return
	}
}

func (handlerApp HandlerApp) Handler() *http.ServeMux {

	handlers := http.NewServeMux()

	handlers.HandleFunc("/", HomeHandler)
	handlers.HandleFunc("/signup", SignupHandler)

	return handlers
}
