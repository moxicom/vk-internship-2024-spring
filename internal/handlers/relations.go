package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/utils"
)

// AddRelation handles the HTTP request to add a relation between a movie and an actor.
//
// @Summary Add relation
// @Description Add a relation between a movie and an actor.
// @Tags Relations
// @Accept json
// @Produce json
// @Param relation body models.RelationMoviesActors true "Relation object to add"
// @Security BasicAuth
// @Success 200 {string} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /relations/ [post]
func (h *Handler) AddRelation(w http.ResponseWriter, r *http.Request) {
	var rel models.RelationMoviesActors
	err := json.NewDecoder(r.Body).Decode(&rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonParseErr, http.StatusBadRequest)
		return
	}
	// Validate JSON body
	validate := validator.New()
	if err := validate.Struct(rel); err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate strings
	err = utils.ValidateRelationParams(rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.AddRelation(rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete by query parameter
// DeleteRelation handles the HTTP request to delete a relation between a movie and an actor.
//
// @Summary Delete relation
// @Description Delete a relation between a movie and an actor.
// @Tags Relations
// @Accept json
// @Produce json
// @Param relation body models.RelationMoviesActors true "Relation object to delete"
// @Security BasicAuth
// @Success 200 {string} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /relations/ [delete]
func (h *Handler) DeleteRelation(w http.ResponseWriter, r *http.Request) {
	// Get rel params
	var rel models.RelationMoviesActors
	err := json.NewDecoder(r.Body).Decode(&rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonParseErr, http.StatusBadRequest)
		return
	}
	// Validate JSON body
	validate := validator.New()
	if err := validate.Struct(rel); err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate strings
	err = utils.ValidateRelationParams(rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Delete relation
	err = h.service.DeleteRelation(rel)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
