package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/hekimapro/task-vault/web/templates/views"
	"github.com/hekimapro/utils/database"
	"github.com/hekimapro/utils/helpers"
	"github.com/hekimapro/utils/log"
	"github.com/hekimapro/utils/server"
)

func main() {

	// database connection
	_, err := database.ConnectToDatabase()
	if err != nil {
		log.Error(err.Error())
	}

	// server context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		views.Index().Render(request.Context(), response)

		// views.Render(request.Context(), response)
	})

	// serving static files
	staticPath := filepath.Join(".", "web", "static")
	fileServer := http.FileServer(http.Dir(staticPath))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// starting server
	server.StartServer(
		ctx,
		mux,
		helpers.GetENVValue("port"),
		helpers.GetENVValue("ssl key path"),
		helpers.GetENVValue("ssl cert path"),
	)
}
