package handlers

import (
	"net/http"
)

func (h *Handler) IsAdminAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, password, ok := r.BasicAuth()
	if ok {

		// TODO: Get from database here
		match, err := h.service.Users.CheckUser(username, password, true)
		if err != nil {
			h.log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if match {
			next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func (h *Handler) IsBasicUserAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, password, ok := r.BasicAuth()
	if ok {
		// You can insert something more here
		match, err := h.service.Users.CheckUser(username, password, false)
		if err != nil {
			h.log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if match {
			next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
