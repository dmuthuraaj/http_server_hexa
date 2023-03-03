package main

import (
	"ex/internal/adapters/left/httpServer"
	"ex/internal/app/routers"
)

func main() {
	// Handler App Adapter
	handlerApp := routers.NewAdapter()

	// Framework Adapters
	// Left httpServer
	httpServer := httpServer.NewAdapter(handlerApp)
	httpServer.Run()
}
