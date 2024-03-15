package handlers

import "net/http"

func (h *handler) getMovies(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get movies request")
}

func (h *handler) getMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get movie request")
}

func (h *handler) addMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("add movie request")
}

func (h *handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("update movie request")
}

func (h *handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("delete movie request")
}
