# syntax=docker/dockerfile:1


#  docker build . -t europe-west3-docker.pkg.dev/mojeico/client/server-grpc:v2
#  docker push europe-west3-docker.pkg.dev/mojeico/client/server-grpc:v2


#  docker build . -t server-grpc:v2
#  docker push server-grpc:v2


FROM golang:1.17-alpine


WORKDIR /app

COPY go.mod go.sum proto ./



RUN go mod download

COPY . .

# MAC OS M1 APPLE  -----RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o main ./cmd/app/main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo  -o main ./main.go

RUN chmod +x ./main


EXPOSE 50051

CMD ["./main"]
