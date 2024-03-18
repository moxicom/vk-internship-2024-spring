package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/utils"
)

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
