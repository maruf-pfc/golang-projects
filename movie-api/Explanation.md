# Movie API in Go

## Package Declaration

```go
package main
```

- This line declares the package name. In Go, the `main` package is special because it defines a standalone executable program. The `main` function within this package is the entry point of the program.

## Imports

```go
    import (
        "encoding/json"
        "fmt"
        "log"
        "math/rand"
        "net/http"
        "strconv"

        "github.com/gorilla/mux"
    )
```

- This block imports necessary packages:
  - `encoding/json`: Provides functions to encode and decode JSON data.
  - `fmt`: Implements formatted I/O operations.
  - `log`: Provides simple logging functions.
  - `math/rand`: Used to generate random numbers.
  - `net/http`: Contains HTTP client and server implementations.
  - `strconv`: Provides functions to convert strings to other types and vice versa.
  - `github.com/gorilla/mux`: A third-party package for routing HTTP requests.

## Struct Definitions

```go
    type Movie struct {
        ID       string   `json:"id"`
        Isbn     string   `json:"isbn"`
        Title    string   `json:"title"`
        Director *Director `json:"director"`
    }
```

- Defines a `Movie` struct with fields for ID, ISBN, title, and a pointer to a `Director` struct. The struct tags (e.g., `json:"id"`) specify how the fields should be encoded/decoded in JSON.

```go
    type Director struct {
        FirstName string `json:"firstname"`
        LastName  string `json:"lastname"`
    }
```

- Defines a `Director` struct with fields for the first and last name, also with JSON struct tags.

## Global Variables

```go
    var movies []Movie
```

- Declares a slice of `Movie` structs to store the list of movies.

## Handler Functions

### Get All Movies

```go
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
```

- `getMovies` is a handler function that writes all movies to the response in JSON format and also prints them in a formatted way.

### Get a Single Movie

```go
    func getMovie(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)
        for _, item := range movies {
            if item.ID == params["id"] {
                json.NewEncoder(w).Encode(item)
                return
            }
        }
        json.NewEncoder(w).Encode(&Movie{})
    }
```

- `getMovie` retrieves a single movie by ID from the URL parameters and writes it to the response. If not found, it returns an empty movie.

### Create a Movie

```go
    func createMovie(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        var movie Movie
        _ = json.NewDecoder(r.Body).Decode(&movie)
        movie.ID = strconv.Itoa(rand.Intn(1000000))
        movies = append(movies, movie)
        json.NewEncoder(w).Encode(movie)
    }
```

- `createMovie` decodes a movie from the request body, assigns it a random ID, appends it to the movies slice, and writes the new movie to the response.

### Update a Movie

```go
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
```

- `updateMovie` finds a movie by ID, removes it, decodes the updated movie from the request body, and appends it back to the slice with the same ID.

### Delete a Movie

```go
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
```

- `deleteMovie` removes a movie by ID from the slice and returns the updated list of movies.

## Main Function

```go
    func main() {
        r := mux.NewRouter()

        movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Spider Man", Director: &Director{FirstName: "John", LastName: "Doe"}})
        movies = append(movies, Movie{ID: "2", Isbn: "654321", Title: "Super Man", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
        movies = append(movies, Movie{ID: "3", Isbn: "987654", Title: "Ant Man", Director: &Director{FirstName: "Jane", LastName: "Doe"}})
        movies = append(movies, Movie{ID: "4", Isbn: "456789", Title: "Iron Man", Director: &Director{FirstName: "Tom", LastName: "Hanks"}})

        r.HandleFunc("/api/movies", getMovies).Methods("GET")
        r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
        r.HandleFunc("/api/movies", createMovie).Methods("POST")
        r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
        r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

        fmt.Println("Server started at port 8080")
        fmt.Println("Go to http://localhost:8080/api/movies")
        log.Fatal(http.ListenAndServe(":8080", r))
    }
```

- `main` initializes the router, adds some sample movies to the slice, sets up the HTTP routes and handlers, and starts the server on port 8080. The server listens for incoming HTTP requests and routes them to the appropriate handler functions.

This code provides a basic RESTful API for managing movies, allowing clients to perform CRUD operations (Create, Read, Update, Delete) on the movie data.
