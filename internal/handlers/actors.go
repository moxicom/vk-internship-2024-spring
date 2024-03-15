package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
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
		h.log.Error(err.Error())
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func (h *handler) getActor(w http.ResponseWriter, r *http.Request, actorId int) {
	h.log.Info("get actor request", "actor_id", actorId)
}

func (h *handler) addActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
	// Decode JSON body
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return
	}
	// Validate JSON body
	validate := validator.New()
	if err := validate.Struct(actor); err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := time.Parse("2006-01-02", actor.BirthDay)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert data
	actorId, err := h.service.Actors.AddActor(actor)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.log.Debug("inserted id", "id", actorId)
	response := struct {
		ID int `json:"id"`
	}{actorId}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *handler) updateActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
}

func (h *handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
}
