package handlers

import (
	"log/slog"
	"net/http"
	"time"

	_ "github.com/moxicom/vk-internship-2024-spring/docs"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
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
	mux.HandleFunc("/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		),
	)
	mux.HandleFunc("/actors/", h.ActorsMainHandler)
	mux.HandleFunc("/movies/", h.MoviesMainHandler)
	mux.HandleFunc("/relations/", h.RelationsMainHandler)
	return mux
}

func (h *Handler) ActorsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.IsBasicUserAuth(w, r, h.GetActorsControler)
	default:
		h.IsAdminAuth(w, r, func(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) MoviesMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.IsBasicUserAuth(w, r, h.GetMoviesController)
	default:
		h.IsAdminAuth(w, r, func(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) RelationsMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.IsAdminAuth(w, r, h.DeleteRelation)
	case http.MethodPost:
		h.IsAdminAuth(w, r, h.AddRelation)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
