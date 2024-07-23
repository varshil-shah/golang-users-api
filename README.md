### GoLang User API using GoFiber

Design a basic web API for user which performs CURD (Create, Update, Read and Delete) operations. Libraries used for building API are GoFiber and MongoDB.

#### API Endpoints

- `GET /users` - Get all users
- `GET /users/{id}` - Get user by id
- `POST /users` - Create a new user
- `PUT /users/{id}` - Update user by id
- `DELETE /users/{id}` - Delete user by id

#### User Model

```go
type User struct {
  ID        string `json:"id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
}
```