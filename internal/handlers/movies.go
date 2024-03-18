package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/utils"
)

var moviesPath = "/movies/"

func (h *Handler) GetMoviesController(w http.ResponseWriter, r *http.Request) {
	idPath := r.URL.Path[len(moviesPath):]
	if len(idPath) == 0 {
		// get all movies
		h.GetMovies(w, r)
	} else {
		// get movie by id
		movieId, err := getIdByPrefix(idPath)
		if err != nil {
			h.log.Error(err.Error())
			http.Error(w, "Invalid movie id", http.StatusBadRequest)
			return
		}
		h.GetMovie(w, r, movieId)
	}
}

// QUERY PARAMETERS: SORT, ORDER, NAME, ACTOR_NAME
// GetMovies handles the HTTP request to retrieve movies with optional sorting and filtering.
//
// @Summary Retrieve movies
// @Description Retrieve a list of movies with optional sorting and filtering.
// @Tags Movies
// @Accept json
// @Produce json
// @Param sort query string false "Sort parameter: valid values are 'name', 'date', 'rating' (default: 'rating')"
// @Param order query string false "Order parameter: valid values are 'asc', 'desc' (default: 'desc')"
// @Param movie_name query string false "Search movies by name (optional)"
// @Param actor_name query string false "Search movies by actor name (optional)"
// @Security BasicAuth
// @Success 200 {array} []models.MovieActors "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movies/ [get]
func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	h.log.Info("get movies request")
	// Get sort params
	var sort models.SortParams
	sort.Sort = r.URL.Query().Get("sort")
	sort.Order = r.URL.Query().Get("order")
	err := utils.ValidateSortParams(sort)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get search params
	var search models.SearchParams
	search.MovieName = r.URL.Query().Get("movie_name")
	search.ActorName = r.URL.Query().Get("actor_name")
	search = utils.ProcessSearchParams(search)
	// Get movies
	movies, err := h.service.Movies.GetMovies(sort, search)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Failed to get movies", http.StatusInternalServerError)
		return
	}
	if len(movies) == 0 {
		w.Write([]byte("[]"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}
}

// GetMovie handles the HTTP request to retrieve a specific movie by ID.
//
// @Summary Retrieve movie
// @Description Retrieve a specific movie by ID.
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID to retrieve"
// @Security BasicAuth
// @Success 200 {object} models.MovieActors "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movies/{id} [get]
func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request, movieId int) {
	h.log.Info("get movie request", "movie_id", movieId)
	movie, err := h.service.Movies.GetMovie(movieId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Failed to get movie", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}
}

// AddMovie handles the HTTP request to add a new movie to the database.
//
// @Summary Add movie
// @Description Add a new movie to the database.
// @Tags Movies
// @Accept json
// @Produce json
// @Param movie body models.Movie true "Movie object to add"
// @Security BasicAuth
// @Success 200 {object} IdResponse "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movies/ [post]
func (h *Handler) AddMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("add movie request")
	// Decode JSON body
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, JsonParseErr, http.StatusBadRequest)
		return
	}
	// Validate JSON body
	validate := validator.New()
	if err := validate.Struct(movie); err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate time string length
	if movie.Date == "" {
		h.log.Error("movie date is required")
		http.Error(w, "movie date is required", http.StatusBadRequest)
		return
	}

	// Validate movie
	err := utils.ValidateMovie(movie)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Insert data
	movieId, err := h.service.Movies.AddMovie(movie)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.log.Debug("inserted id", "id", movieId)
	response := struct {
		ID int `json:"id"`
	}{movieId}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonEncodeErr, http.StatusInternalServerError)
		return
	}

}

// Update movie by movie id
// UpdateMovie handles the HTTP request to update an existing movie by its ID.
//
// @Summary Update movie
// @Description Update an existing movie by its ID.
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID to update"
// @Param movie body models.Movie true "Updated movie object"
// @Security BasicAuth
// @Success 200 {string} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movies/{id} [put]
func (h *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("update movie request")
	// Check movie id existance in URL
	idPath := r.URL.Path[len(moviesPath):]
	if len(idPath) == 0 {
		h.log.Error("Invalid movie id")
		http.Error(w, "Invalid movie id", http.StatusBadRequest)
		return
	}
	// Get movie id
	movieId, err := getIdByPrefix(idPath)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid movie id", http.StatusBadRequest)
		return
	}
	// Decode JSON body
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		h.log.Error(err.Error())
		http.Error(w, JsonParseErr, http.StatusBadRequest)
		return
	}
	// Validate movie
	err = utils.ValidateMovie(movie)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update data
	err = h.service.Movies.UpdateMovie(movieId, movie)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteMovie handles the HTTP request to delete a movie by its ID.
//
// @Summary Delete movie
// @Description Delete a movie by its ID.
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID to delete"
// @Security BasicAuth
// @Success 200 {string} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /movies/{id} [delete]
func (h *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	h.log.Info("delete movie request")
	// Check movie id existance in URL
	idPath := r.URL.Path[len(moviesPath):]
	if len(idPath) == 0 {
		h.log.Error("Invalid movie id")
		http.Error(w, "Invalid movie id", http.StatusBadRequest)
		return
	}
	// Get movie id
	movieId, err := getIdByPrefix(idPath)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid movie id", http.StatusBadRequest)
		return
	}
	h.log.Debug("delete movie", "movie_id", movieId)
	// Delete movie by id
	err = h.service.Movies.DeleteMovie(movieId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
