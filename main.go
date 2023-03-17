package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", isbn: "438228", Title: "Movie two", Director: &Director{Firstname: "Shasha", Lastname: "Yang"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("PUT")
	r.Handlefunc("/movies{id}", updateMovies).Methods("POST")
	r.HandleFunc("/movies{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Server is starting at PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000"))
}
