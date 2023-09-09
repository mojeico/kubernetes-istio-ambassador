# syntax=docker/dockerfile:1

FROM golang:1.17-alpine


#  docker build --platform linux/amd64 . -t  mojeico/golang-client:v4
#  docker push  mojeico/golang-client:v4


WORKDIR /app

COPY go.mod go.sum proto ./



RUN go mod download

COPY . .

# MAC OS M1 APPLE  -----RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o main ./cmd/app/main.go

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo  -o main ./main.go


RUN CGO_ENABLED=0 go build -a -installsuffix cgo  -o main ./main.go


RUN chmod +x ./main


EXPOSE 3000

CMD ["./main"]


