
### Build Design-First APIs with Go

Rapid prototyping of APIs from design to implementation.

#### 1. Design APIs with Stoplight Studio https://stoplight.io

        a. Refer to OpenAPI specs in apis/users.yaml
        b. API endpoints
            -  POST /users
            -  GET /users/{userId}
            -  PATCH /users/{userId}

#### 2. Generate Gin Server boilerplate from OpenAPI specs 

        a. Use "github.com/deepmap/oapi-codegen" to generate boilerplate code
        b. Example
 
           go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=Users --generate types,gin -o users/users.gen.go apis/users.yaml

#### 3. Server Implementation

        a. Add handler functions in users/users.go
        b. Create a Gin router and register the handler in main.go

#### 4. Run API server

        go run .












