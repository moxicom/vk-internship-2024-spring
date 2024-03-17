package handlers

import "net/http"

func (h *Handler) getMovies(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get movies request")
}

func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get movie request")
}

func (h *Handler) addMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("add movie request")
}

func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("update movie request")
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("delete movie request")
}
