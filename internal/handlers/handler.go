package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/moxicom/vk-internship-2024-spring/internal/service"
)

type handler struct {
	log     *slog.Logger
	service *service.Service
}

func Run(logger *slog.Logger, s *service.Service) error {
	handler := &handler{logger, s}
	mux := handler.initRoutes()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	return server.ListenAndServe()
}

func (h *handler) initRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", h.ping)
	mux.HandleFunc("/actors/", h.actorsMainHandler)
	mux.HandleFunc("/movies/", h.moviesMainHandler)
	return mux
}

func (h *handler) actorsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getActorsControler(w, r)
	default:
		h.withMiddleware(w, r, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				h.addActor(w, r)
			case http.MethodPut:
				h.updateActor(w, r)
			case http.MethodDelete:
				h.deleteActor(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}

func (h *handler) moviesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getMovies(w, r)
	default:
		h.withMiddleware(w, r, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				h.addMovie(w, r)
			case http.MethodPut:
				h.updateMovie(w, r)
			case http.MethodDelete:
				h.deleteMovie(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}

func (h *handler) withMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h.log.Info("middleware")
	next(w, r)
}
