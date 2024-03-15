package handlers

import (
	"encoding/json"
	"net/http"
)

var actorsPath = "/actors/"

func (h *handler) getActorsControler(w http.ResponseWriter, r *http.Request) {
	idPath := r.URL.Path[len(actorsPath):]
	if len(idPath) == 0 {
		// get all actors
		h.getActors(w, r)
	} else {
		// get actor by id
		actorId, err := getIdByPrefix(idPath)
		if err != nil {
			http.Error(w, "Invalid actor id", http.StatusBadRequest)
			return
		}
		h.getActor(w, r, actorId)
	}
}

func (h *handler) getActors(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
	actors, err := h.service.GetActors()
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if len(actors) == 0 {
		w.Write([]byte("[]"))
		return
	}

	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) getActor(w http.ResponseWriter, r *http.Request, actorId int) {
	h.log.Info("get actor request", "actor_id", actorId)
}

func (h *handler) addActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")

}

func (h *handler) updateActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
}

func (h *handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
}
