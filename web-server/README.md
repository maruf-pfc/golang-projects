# Simple Go Web Server

This is a simple web server written in [Go](https://golang.org/doc/install). It serves static files and handles basic form submissions. The server listens on port 8080 and provides two main functionalities: serving a static HTML form and responding to a "hello" endpoint.

## Features

- **Static File Serving**: Serves static files from the `./static` directory.
- **Form Handling**: Processes form submissions via the `/form` endpoint.
- **Hello Endpoint**: Responds with a simple greeting at the `/hello` endpoint.

## Endpoints

### `/`

- **Method**: GET
- **Description**: Serves static files located in the `./static` directory. Ensure that your `form.html` file is placed in this directory to be served correctly.

### `/form`

- **Method**: GET, POST
- **GET**: Serves the `form.html` file which contains the form for user input.
- **POST**: Processes the submitted form data. It expects two fields: `name` and `address`. The server will print these values to the console and also display them back to the user in the browser.

### `/hello`

- **Method**: GET
- **Description**: Responds with a simple "hello!" message. If accessed with any method other than GET, or if the path is incorrect, it returns a 404 error.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

### Running the Server

1. Clone the repository or copy the Go file to your local machine.
2. Ensure you have a `form.html` file in a directory named `static` in the same location as your Go file.
3. Open a terminal and navigate to the directory containing the Go file.
4. Run the server using the command:

   ```bash
   go run main.go
   ```

5. Open your web browser and go to `http://localhost:8080` to access the static files or `http://localhost:8080/form` to access the form.

## Example `form.html`

Here is a simple example of what your `form.html` might look like:

![image](https://github.com/user-attachments/assets/68d47461-3d4d-49c7-b6da-ba9518b998af)


## Notes

- Ensure that the `static` directory is correctly set up with the necessary HTML files.
- The server logs form submissions to the console for easy debugging and verification.

## License

This project is open-source and available under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute as needed.
