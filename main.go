package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContecreateMoviesnt-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // this will return the remaining movies
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie two", Director: &Director{Firstname: "Shasha", Lastname: "Yang"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("PUT")
	r.HandleFunc("/movies{id}", updateMovies).Methods("POST")
	r.HandleFunc("/movies{id}", deleteMovies).Methods("DELETE")
	fmt.Printf("Server is starting at PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
