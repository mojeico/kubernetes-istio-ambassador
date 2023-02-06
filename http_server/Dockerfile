# syntax=docker/dockerfile:1


#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/http-server:v1
#  docker push europe-west3-docker.pkg.dev/mojeico/client/http-server:v1

#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/http-server:v2
#  docker push europe-west3-docker.pkg.dev/mojeico/client/http-server:v2

#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/http-server:v3
#  docker push europe-west3-docker.pkg.dev/mojeico/client/http-server:v3

#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/http-server:v0-mirroring
#  docker push europe-west3-docker.pkg.dev/mojeico/client/http-server:v0-mirroring


#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/http-server:v0-exec1
#  docker push europe-west3-docker.pkg.dev/mojeico/client/http-server:v0-exec1


FROM golang:1.17

WORKDIR /app

COPY . ./



LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates"
LABEL alias="yourtool"


RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /docker-gs-ping

EXPOSE 3001

CMD [ "/docker-gs-ping" ]
