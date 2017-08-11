## Instructions

> A simple RESTful API microservice in Go

### App Dependency

Make sure you have the following installed:

1. MySQL, and running on local machine
2. Go Language (1.8 or above)

### Install App
1. Clone the application
2. Install Go package dependency

  go get github.com/gin-gonic/gin
  go get gopkg.in/gorp.v1
  go get github.com/go-sql-driver/mysql

3. Log into MySQL, create a `instructions` database

### Testing App
Run:

  go test

### Run App
1. Start Go app locally

  go run main.go

2. Play with instruction endpoints

* Create a new instruction

  curl -i -X POST -H "Content-Type: application/json" -d "{ \"event_status\": \"83\", \"event_name\": \"100\" }" http://localhost:8080/api/v1/instructions

* Show an existing instruction

  curl -i http://localhost:8080/api/v1/Instructions/1


* Show all existing instructions

  curl -i http://localhost:8080/api/v1/instructions

* Delete an existing instructions

  curl -i -X DELETE http://localhost:8080/api/v1/instructions/1

* Update an existing instructions

  curl -i -X PUT -H "Content-Type: application/json" -d "{ \"event_status\": \"83\", \"event_name\": \"100\" }" http://localhost:8080/api/v1/instructions/1
