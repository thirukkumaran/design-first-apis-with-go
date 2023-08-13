
### Build Design-First APIs with Go

Rapid prototyping of API server from design to implementation.

#### 1. Design APIs with Stoplight Studio https://stoplight.io

1. Refer OpenAPI sepcs in apis/users.yaml
2. API endpoints
    -  POST /users
    -  GET /users/{userId}
    -  PATCH /users/{userId}

#### 2. Generate Gin Server boileplate from OpenAPI specs 

1. Features of Gin web framework https://gin-gonic.com/docs/introduction/
2. Use "github.com/deepmap/oapi-codegen" to generate boilerplate code
 
       go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=Users --generate types,gin -o users/users.gen.go apis/users.yaml

#### 3. Server Implementation

1. Add business logic in handler functions in users/users.go
2. Create Gin router and register the handler in main.go

#### 4. Run API server

go run .












