# Movie API

This is a simple RESTful API for managing a collection of movies. It is built using Go and the Gorilla Mux router.

## Features

- **Get all movies**: Retrieve a list of all movies.
- **Get a single movie**: Retrieve details of a specific movie by ID.
- **Create a new movie**: Add a new movie to the collection.
- **Update a movie**: Modify details of an existing movie.
- **Delete a movie**: Remove a movie from the collection.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Gorilla Mux package

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/maruf-pfc/golang-projects.git
   cd movie-api
   ```

2. Install the Gorilla Mux package:

   ```bash
   go get -u github.com/gorilla/mux
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

The server will start on port 8080.

### API Endpoints

- **GET /api/movies**: Retrieve all movies.
- **GET /api/movies/{id}**: Retrieve a movie by ID.
- **POST /api/movies**: Create a new movie. Requires a JSON body with `isbn`, `title`, and `director` fields.
- **PUT /api/movies/{id}**: Update an existing movie by ID. Requires a JSON body with `isbn`, `title`, and `director` fields.
- **DELETE /api/movies/{id}**: Delete a movie by ID.

### Example Movie JSON

```json
{
  "isbn": "123456",
  "title": "Spider Man",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
