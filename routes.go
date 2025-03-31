package main

import (
	"simple_api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.StripSlashes, LogHandler)

	// Beatmap Routes
	r.Route("/beatmap", func(r chi.Router) {
		r.Get("/", handler.GetBeatmaps)
		r.Post("/", handler.InsertBeatmap)
	})

	return r
}
