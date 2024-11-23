# REST API with Go and Gin

This is a sample REST API built using Go, with the **Gin** web framework for handling routes and HTTP requests. The API allows the creation and management of events, as well as user registration and authentication via JWT tokens.

## Technologies Used

- **Go (Golang)**: The programming language used to develop the API.
- **Gin**: A high-performance web framework for Go used to manage routes and HTTP requests.
- **SQLite**: The database used to store user, event, and registration information.
- **JWT (JSON Web Tokens)**: Used for user authentication through token generation and validation.

## External Dependencies

The following external libraries are used in the application:

- **Gin**: Web framework for Go to handle HTTP routes and middleware.
  - Installation: `github.com/gin-gonic/gin`
- **SQLite**: SQLite database driver for Go.
  - Installation: `modernc.org/sqlite`
- **Godotenv**: Loads environment variables from a `.env` file.
  - Installation: `github.com/joho/godotenv`
- **JWT-Go**: Library for creating and verifying JWT tokens.
  - Installation: `github.com/golang-jwt/jwt/v5`
- **Crypto**: Library that holds cryptography packages.
  - Installations `golang.org/x/crypto`

## Available Endpoints

### Public Routes

- `POST /signup`: Create a new user. Data should be sent in the request body in JSON format (email, password).
- `POST /login`: Log in with an existing user and receive a JWT token. Data should be sent in the request body in JSON format (email, password).

### Authenticated Routes (require JWT)

- `GET /events`: Get a list of all events.
- `GET /events/:id`: Get the details of a specific event by ID.
- `POST /events`: Create a new event (requires authentication).
- `PUT /events/:id`: Update an existing event (requires authentication).
- `DELETE /events/:id`: Delete an event (requires authentication).
- `POST /events/:id/register`: Register a user for an event (requires authentication).
- `DELETE /events/:id/register`: Cancel a user's registration for an event (requires authentication).

## How to Run the Application

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-user/rest-api.git
   cd rest-api
   ```
2. **Install dependencies**: Ensure you have Go installed. If not, you can download and install it from here
3. **Set up the .env file**: Create a .env file in the root directory of the project with the necessary environment variables. For example
   ```bash
   JWT_SECRET_KEY=supersecret
   ```
4. **Start the application**: Run the following command to start the development server:
   ```bash
   go run main.go
   ```
   This will start the server on port 8080, and you can begin making requests to the defined endpoints.

## Key Architecture Points

### Authentication and JWT Generation

The application uses JSON Web Tokens (JWT) for authentication. The flow is as follows:

- When a user registers (`/signup`), a new user is created in the database.
- When a user logs in (`/login`), the application generates a JWT token using a secret key defined in the `.env` file (`JWT_SECRET_KEY`).
- This token is returned to the user and must be included in the headers of authenticated requests as part of the authorization: `Authorization: Bearer <token>`.
- Protected routes (such as `POST /events`, `PUT /events/:id`, etc.) require the user to send this token in the request headers. The application verifies the token before processing the request.

### JWT Token Generation

The JWT token is generated with the following code in `utils/jwt.go`. This code uses the `jwt-go` library to generate a token with user details (email, userId) and an expiration time of 2 hours.

### Authentication Middleware

The Authenticate middleware ensures the user is authenticated before accessing protected routes. It checks that the JWT token sent in the request header is valid. This middleware verifies the JWT token, and if it’s valid, it adds the userId to the request context, allowing controllers to access the authenticated user.

## Project Structure

- `/db`: Database handling, including table creation and executing SQL queries.
- `/env`: Manages environment variables using the godotenv package.
- `/middlewares`: Middleware for authenticating protected routes.
- `/routes`: Defines all the API routes, both public and protected.
- `/utils`: Utility functions, such as token generation and verification.

## License

This project is licensed under the MIT License.

```bash
This version of the `README.md` file is in English and covers everything you requested, including the technologies, dependencies, endpoints, setup instructions, key architecture points, and more. Let me know if you need any further changes!
```
