# Movie Database Management REST API
This REST API is designed for managing a movie database. It provides functionalities for adding, updating, deleting, and retrieving information about movies and actors. The API supports both regular users and administrators.

### User Roles:
Administrator: Has full access to all functionalities of the API.
Regular User: Can view movie and actor information and perform searches.

### How to use
- update ```.env``` file. It can be based on ```.env.example``` file
- run docker compose by command ```docker-compose up --build```. It will be create a postgres database and run sql scripts
- run server by command ```go run cmd/main.go``` or ```make run```

Server runs locally on port ```:8080```. To check REST API documentation open link below.
```
http://localhost:8080/swagger/
```