package httpServer

import (
	"ex/internal/ports"
	"log"
	"net/http"
	"os"
)

type ServerAdpater struct {
	router ports.RouterApp
}

func NewAdapter(routera ports.RouterApp) *ServerAdpater {
	return &ServerAdpater{router: routera}
}

func (server ServerAdpater) Run() {
	var httpServerPort string

	httpServerPort = os.Getenv("HTTP_SERVER_PORT")
	if httpServerPort == "" {
		httpServerPort = ":9100"
	}

	s := &http.Server{
		Addr:    httpServerPort,
		Handler: server.router.Handler(),
	}

	log.Printf("Starting server on port %v ...", httpServerPort)
	if err := s.ListenAndServe(); err != nil {
		log.Printf("unable to start listening on port %v error:%v\n", httpServerPort, err)
	} else {
		log.Printf("Server listening on port %v", httpServerPort)
	}
}
