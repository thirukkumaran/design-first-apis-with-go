
### Build Design-First APIs with Go

Rapid prototyping of API server from design to implementation.

#### 1. Design APIs with Stoplight Studio https://stoplight.io

a. Refer OpenAPI sepcs in apis/users.yaml
b. API endpoints
    -  POST /users
    -  GET /users/{userId}
    -  PATCH /users/{userId}

#### 2. Generate Gin Server boileplate from OpenAPI specs 

a. Features of Gin web framework https://gin-gonic.com/docs/introduction/
b. Use "github.com/deepmap/oapi-codegen" to generate boilerplate code
 
       go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=Users --generate types,gin -o users/users.gen.go apis/users.yaml

#### 3. Server Implementation

a. Add business logic in handler functions in users/users.go
b. Create Gin router and register the handler in main.go

#### 4. Run API server

go run .












