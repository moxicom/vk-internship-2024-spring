package handlers

import (
	"net/http"
)

func (h *Handler) isAdminAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, password, ok := r.BasicAuth()
	if ok {

		// TODO: Get from database here
		expectedUsername := "userAdmin"
		expectedPassword := "password"

		usernameMatch := username == expectedUsername
		passwordMatch := password == expectedPassword

		if usernameMatch && passwordMatch {
			next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func (h *Handler) isBasicUserAuth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	username, password, ok := r.BasicAuth()
	if ok {

		// TODO: Get from database here
		expectedUsername := "user"
		expectedPassword := "password"

		usernameMatch := username == expectedUsername
		passwordMatch := password == expectedPassword

		if usernameMatch && passwordMatch {
			next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
