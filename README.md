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
There are three go files:
  - **server/server.go** -  creates a http server that exercises GET, POST, PUT, and DELETE.
    - homePage() - a homepage that reveals a message only to authorized clients.
    - isAuthorized() - validates the token from incoming requests
  - **server/main_test.go** - contains unit tests to check the functionalities
  - **client/client.go** - a client API to generate a signed JWT to hit the server endpoint.
     - homePage() - a homepage that makes a GET request to the server's homepage, using the generated token
     - GenerateJWT() - a function to generate a token to access the server's homepage

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

After running the server, you can check to the homepage: http://localhost:9000/

You will find the message: "Not Authorized".

After running the client:
``` bash
cd client
go build
go run client.go
```
You can go to port 8001 on the client side: http://localhost:8001/
And you will find a message: "Only Authorized Clients can see this message!" This only works when the client side has the same value in the 'mySigningKey' variable as the server side. Otherwise, it will fail.

Note: In practice, the key should be stored as an environmental variable.
