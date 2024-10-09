package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Init movies var as a slice Movie struct
var movies []Movie

// Get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	fmt.Fprintf(w, "\n\nAll movies:\n")
	for _, movie := range movies {
		fmt.Fprintf(w, "ID: %s\n", movie.ID)
		fmt.Fprintf(w, "ISBN: %s\n", movie.Isbn)
		fmt.Fprintf(w, "Title: %s\n", movie.Title)
		fmt.Fprintf(w, "Director: %s %s\n\n", movie.Director.FirstName, movie.Director.LastName)
	}
}

// Get single movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through movies and find with id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Delete movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// - `movies[:index]`: This creates a new slice that contains all elements of the `movies` slice from the start up to, but not including, the element at the specified `index`. 

// - `movies[index+1:]`: This creates another slice that contains all elements of the `movies` slice starting just after the specified `index` to the end of the slice.

// - `append(movies[:index], movies[index+1:]...)`: The `append` function is used to concatenate these two slices. The `...` is a variadic argument that expands the second slice so that each of its elements is appended individually. This effectively skips the element at the specified `index`, thereby removing it from the original slice.

// The result is a new slice that contains all the elements of the original `movies` slice except for the one at the specified `index`. This operation does not modify the original slice in place but instead creates a new slice with the desired elements.
        
func main() {
	// Init router
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Spider Man", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "654321", Title: "Super Man", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "987654", Title: "Ant Man", Director: &Director{FirstName: "Jane", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "4", Isbn: "456789", Title: "Iron Man", Director: &Director{FirstName: "Tom", LastName: "Hanks"}})

	// Route handles & endpoints
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	// Start server
	fmt.Println("Server started at port 8080")
	fmt.Println("Go to http://localhost:8080/api/movies")
	log.Fatal(http.ListenAndServe(":8080", r))
}