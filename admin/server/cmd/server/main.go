package main

import (
	"github.com/batazor/fsme/admin/server/pkg/handler"
	l "github.com/batazor/fsme/admin/server/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// Logger ==================================================================
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Get configuration =======================================================
	PORT := "6001"

	// Routes ==================================================================
	r := chi.NewRouter()

	// CORS ====================================================================
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//Debug:            true,
	})

	r.Use(cors.Handler)
	r.Use(l.Logger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Recoverer)
	r.NotFound(NotFoundHandler)

	r.Mount("/", handler.Routes())
	logger.Info("Run on port " + PORT)

	// start HTTP-server
	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	return
}
