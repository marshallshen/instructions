## Instructions

> A simple RESTful API microservice in Go

### Prerequisites

1. [Go Language](https://golang.org) 1.8 or later.

2. [MySQL Server](https://hub.docker.com/_/mysql) 5.7.

You can run MySQL in a docker container as follows.

We will set some environment variables that will help us to run our MySQL server container and we will also leverage these environment variables when we run our microservice.

```
export MYSQL_CONTAINER_NAME=mysql_instruct # relabel as appropriate
export MYSQL_ROOT_PASSWORD=root_password # replace as appropriate
export MYSQL_DATABASE=instructions
export MYSQL_USER=instruct
export MYSQL_PASSWORD=instruct_password # replace as appropriate
```

To run our MySQL Server:
```
docker run \
  --name=${MYSQL_CONTAINER_NAME} \
  -d \
  -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
  -e MYSQL_DATABASE=${MYSQL_DATABASE} \
  -e MYSQL_USER=${MYSQL_USER} \
  -e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
  -p 3306:3306 \
  mysql/mysql-server:5.7
```

### Install App
1. Clone the repo.
2. Install Go package dependencies

```
go get github.com/gin-gonic/gin
go get gopkg.in/gorp.v1
go get github.com/go-sql-driver/mysql
```

### Testing App

```
go test
```

### Run App
1. Start Go app locally

```
go run main.go
```

2. Play with instruction endpoints

* Create a new instruction

```
curl \
  -i \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{ "event_status": "83", "event_name": "100" }' \
  http://localhost:8080/api/v1/instructions
```

* Show an existing instruction

```
curl -i http://localhost:8080/api/v1/instructions/1
```

* Show all existing instructions

```
curl -i http://localhost:8080/api/v1/instructions
```

* Delete an existing instructions

```
curl -i -X DELETE http://localhost:8080/api/v1/instructions/1
```

* Update an existing instructions

```
curl \
  -i \
  -X PUT \
  -H "Content-Type: application/json" \
  -d '{ "event_status": "83", "event_name": "100" }' \
  http://localhost:8080/api/v1/instructions/1
```

### Run App with Docker

**Make sure Docker is installed before executing the command below**

```
docker build -t instructions-app . # inside the app directory
```

### Cleanup

1. Press CTRL-C to stop the application.

2. Stop MySQL Server if you started it using Docker.

```
docker stop ${MYSQL_CONTAINER_NAME}
docker rm ${MYSQL_CONTAINER_NAME}
```
