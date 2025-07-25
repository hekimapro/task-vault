package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/hekimapro/task-vault/internal/handlers"
	"github.com/hekimapro/task-vault/internal/repositories"
	"github.com/hekimapro/task-vault/internal/services"
	"github.com/hekimapro/task-vault/web/templates/views"
	"github.com/hekimapro/utils/database"
	"github.com/hekimapro/utils/helpers"
	"github.com/hekimapro/utils/log"
	"github.com/hekimapro/utils/server"
)

func main() {

	// server mux
	mux := http.NewServeMux()

	// database connection
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Error(err.Error())
	}

	// server context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	//handlers
	userHandler := handlers.NewUser(services.NewUser(repositories.NewUser(db)))

	mux.HandleFunc("GET /login", userHandler.LoginView)
	mux.HandleFunc("POST /login-post", userHandler.Login)
	mux.HandleFunc("POST /register-post", userHandler.Create)

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
