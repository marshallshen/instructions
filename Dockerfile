FROM        golang:1.8
MAINTAINER  Marshall Shen <shen.marshall@gmail.com>

ENV     PORT  8080
   
# Setting up working directory
WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container

RUN     go get github.com/tools/godep
RUN     go get github.com/gin-gonic/gin
RUN     go get gopkg.in/gorp.v1
RUN     go get github.com/go-sql-driver/mysql
RUN     go install github.com/tools/godep
RUN     go install github.com/gin-gonic/gin

# Restore godep dependencies
#RUN godep restore

EXPOSE 8080
ENTRYPOINT  ["/usr/local/bin/go"]
CMD     ["run", "main.go"]