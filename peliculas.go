package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
	Length int    `json:"length"`
}

var movies []Movie

func peliculas() {
	http.HandleFunc("/movies", handleMovies)
	http.HandleFunc("/movies/search", handleMovieSearch)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMovies(w, r)
	case http.MethodPost:
		addMovie(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
}

func handleMovieSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	results := searchMovies(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func searchMovies(query string) []Movie {
	var results []Movie

	for _, movie := range movies {
		if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(movie.Genre), strings.ToLower(query)) {
			results = append(results, movie)
		}
	}

	return results
}
