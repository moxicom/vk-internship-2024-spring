package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

var actorsPath = "/actors/"

type IdResponse struct {
	ID int `json:"id"`
}

func (h *Handler) GetActorsControler(w http.ResponseWriter, r *http.Request) {
	idPath := r.URL.Path[len(actorsPath):]
	if len(idPath) == 0 {
		// get all actors
		h.GetActors(w, r)
	} else {
		// get actor by id
		actorId, err := getIdByPrefix(idPath)
		if err != nil {
			http.Error(w, "Invalid actor id", http.StatusBadRequest)
			return
		}
		h.GetActor(w, r, actorId)
	}
}

// GetActors handles the HTTP request to retrieve all actors.
//
// @Summary Retrieve actors
// @Description Retrieve a list of all actors.
// @Tags Actors
// @Accept json
// @Security BasicAuth
// @Produce json
// @Success 200 {array} []models.ActorFilms "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /actors/ [get]
func (h *Handler) GetActors(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get actors request")
	actors, err := h.service.Actors.GetActors()
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	if len(actors) == 0 {
		w.Write([]byte("[]"))
		return
	}
	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}
}

// @Summary Retrieve actors
// @Description Retrieve a list of all actors or a specific actor by ID.
// @Tags Actors
// @Param id path string false "Actor ID to retrieve (optional)"
// @Accept json
// @Security BasicAuth
// @Produce json
// @Success 200 {array} models.ActorFilms "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /actors/{id} [get]
func (h *Handler) GetActor(w http.ResponseWriter, r *http.Request, actorId int) {
	h.log.Info("get actor request", "actor_id", actorId)
	actor, err := h.service.Actors.GetActor(actorId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if actor.ID == 0 && actor.Name == "" && actor.BirthDay == "" && actor.Gender == "" {
		w.Write([]byte("{}"))
		return
	}
	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}
}

// AddActor handles the HTTP request to add a new actor.
//
// @Summary Add actor
// @Description Add a new actor to the database.
// @Tags Actors
// @Accept json
// @Security BasicAuth
// @Param actor body models.Actor true "Actor object to add"
// @Produce json
// @Success 200 {object} IdResponse "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /actors/ [post]
func (h *Handler) AddActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("add actor request")
	// Decode JSON body
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, JsonParseErr, http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}
}

// UpdateActor handles the HTTP request to update an existing actor.
//
// @Summary Update actor
// @Description Update an existing actor in the database.
// @Tags Actors
// @Accept json
// @Security BasicAuth
// @Param id path string true "Actor ID to update"
// @Param actor body models.Actor true "Actor object to update"
// @Produce json
// @Success 200 "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /actors/{id} [put]
func (h *Handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("update actors request")
	// Check actor id existance in URL
	idPath := r.URL.Path[len(actorsPath):]
	if len(idPath) == 0 {
		h.log.Error("Invalid actor id")
		http.Error(w, "Invalid actor id", http.StatusBadRequest)
		return
	}
	// Get actor id
	actorId, err := getIdByPrefix(idPath)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid actor id", http.StatusBadRequest)
		return
	}
	// Decode JSON body
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonParseErr, http.StatusBadRequest)
		return
	}
	// Validate time if it is not empty
	if actor.BirthDay != "" {
		_, err := time.Parse("2006-01-02", actor.BirthDay)
		if err != nil {
			h.log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	// Update data
	err = h.service.Actors.UpdateActor(actorId, actor)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteActor handles the HTTP request to delete an existing actor.
//
// @Summary Delete actor
// @Description Delete an existing actor from the database.
// @Tags Actors
// @Accept json
// @Security BasicAuth
// @Param id path string true "Actor ID to delete"
// @Produce json
// @Success 200 "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /actors/{id} [delete]
func (h *Handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	h.log.Info("delete actors request")
	idPath := r.URL.Path[len(actorsPath):]
	if len(idPath) == 0 {
		err := errors.New("no actor id in URL")
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	actorId, err := getIdByPrefix(idPath)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid actor id", http.StatusBadRequest)
		return
	}
	// delete actor by id
	err = h.service.Actors.DeleteActor(actorId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
