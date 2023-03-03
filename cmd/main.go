package main

import (
	"ex/internal/adapters/left/httpServer"
	"ex/internal/app/handlers"
)

func main() {
	// Handler App Adapter
	handlerApp := handlers.NewAdapter()

	// Framework Adapters
	// Left httpServer
	httpServer := httpServer.NewAdapter(handlerApp)
	httpServer.Run()
}
