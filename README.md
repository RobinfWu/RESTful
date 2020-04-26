# RESTful API | Unit Testing | Securing Endpoints

## Quick Start
``` bash
# Install mux router
go get -u github.com/gorilla/mux
```
``` bash
cd server
go build
go run server.go
```
After running the server, you can access: http://localhost:9000/patients

Run main_test.go
``` bash
go test
```
There are two main folders:
  - **server/server.go** -  creates a http server that exercises GET, POST, PUT, and DELETE.
  - **server/main_test.go** - contains unit tests to check the functionalities
  - **client/client.go** - a client API to generate a signed JWT to hit the server endpoint.

## Endpoints

### Get All Patients
``` bash
GET patients
```
### Get Single Patient
``` bash
GET patients/{id}
```

### Delete Patient
``` bash
DELETE patients{id}
```

### Create Patient
``` bash
POST patients

# Request sample
# {
#   "ID":"4",
#   "firstname":"John",
#   "lastname":"Doe",
#   "address":"7 Address ST",
#   "Doctor":{"firstname":"Mary",  "lastname":"Sue"}
# }
```
### Update Patient
``` bash
PUT patients/{id}

# Request sample
# {
#   "ID":"4",
#   "firstname":"John",
#   "lastname":"Doe",
#   "address":"7 Address ST",
#   "Doctor":{"firstname":"Mary",  "lastname":"Sue"}
# }
```
### Bonus: Securing Endpoints
Endpoints can be secured using JWT (JSON Web Tokens) Authentication. When a client requests data to the server, it generates a signed JWT to be part of its request. The server reads the JWT and validates the token.
