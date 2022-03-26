package internal

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()

	// base middleware stack from https://github.com/go-chi/chi
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// handle spa routes
	r.Get("/*", app.handleView)

	// static routes
	// dev mode
	// todo: figure out prod
	fileServer := http.FileServer(http.Dir("./web/view/src/"))
	r.Handle("/src*", http.StripPrefix("/src", fileServer))

	return r
}
