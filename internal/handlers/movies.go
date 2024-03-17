package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/utils"
)

var moviesPath = "/movies/"

func (h *Handler) getMoviesController(w http.ResponseWriter, r *http.Request) {
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
	// Validate time string
	_, err := time.Parse("2006-01-02", movie.Date)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate rating
	if *movie.Rating < 0 || *movie.Rating > 10 {
		h.log.Error(MovieRatingErr)
		http.Error(w, MovieRatingErr, http.StatusBadRequest)
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
	// Validate time string if it is not empty
	if movie.Date != "" {
		_, err := time.Parse("2006-01-02", movie.Date)
		if err != nil {
			h.log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	// Validate rating if it is not empty
	if movie.Rating != nil {
		if *movie.Rating < 0 || *movie.Rating > 10 {
			h.log.Error(MovieRatingErr)
			http.Error(w, MovieRatingErr, http.StatusBadRequest)
			return
		}
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
	// Delete movie by id
	err = h.service.Movies.DeleteMovie(movieId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
