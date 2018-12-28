package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/reyhanfahlevi/soap-absence/api/http/absence"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You know... for nakama attendance"))
}

func handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Get("/", root)
	r.Route("/device", func(r chi.Router) {
		r.Post("/save", absence.HandlerAddNewDevice)
	})

	return r
}
