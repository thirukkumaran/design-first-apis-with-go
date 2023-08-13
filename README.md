
### Build Design-First APIs with Go

Rapid prototyping of API server from design to implementation.

### Design APIs with Stoplight Studio https://stoplight.io
1. Refer OpenAPI sepcs in apis/users.yaml
2. API endpoints
    -  POST /users
    -  GET /users/{userId}
    -  PATCH /users/{userId}

#### Generate Gin Server boileplate from OpenAPI specs 

1. Featurs of Gin web framework https://gin-gonic.com/docs/introduction/
2. Use github.com/deepmap/oapi-codegen to generate boilerplate code
 
       go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=Users --generate types,gin -o users/users.gen.go apis/users.yaml

### Server Implementation
1. Add business logic in handler functions in users/users.go
2. Create Gin router and register handler in main.go

### Run API server
go run .












