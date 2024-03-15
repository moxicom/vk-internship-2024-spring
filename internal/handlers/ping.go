package handlers

import "net/http"

func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ping request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
