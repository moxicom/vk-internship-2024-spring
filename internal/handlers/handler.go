package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/moxicom/vk-internship-2024-spring/internal/service"
)

type Handler struct {
	log     *slog.Logger
	service *service.Service
}

func NewHandler(log *slog.Logger, s *service.Service) *Handler {
	return &Handler{log, s}
}

func Run(logger *slog.Logger, s *service.Service) error {
	handler := NewHandler(logger, s)
	mux := handler.InitRoutes()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	return server.ListenAndServe()
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/actors/", h.actorsMainHandler)
	mux.HandleFunc("/movies/", h.moviesMainHandler)
	mux.HandleFunc("/relations/", h.relationsMainHandler)
	return mux
}

func (h *Handler) actorsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetActorsControler(w, r)
	default:
		h.withMiddleware(w, r, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				h.AddActor(w, r)
			case http.MethodPut:
				h.UpdateActor(w, r)
			case http.MethodDelete:
				h.DeleteActor(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}

func (h *Handler) moviesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetMoviesController(w, r)
	default:
		h.withMiddleware(w, r, func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				h.AddMovie(w, r)
			case http.MethodPut:
				h.UpdateMovie(w, r)
			case http.MethodDelete:
				h.DeleteMovie(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}
}

func (h *Handler) relationsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.DeleteRelation(w, r)
	case http.MethodPost:
		h.AddRelation(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) withMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h.log.Info("middleware")
	next(w, r)
}
