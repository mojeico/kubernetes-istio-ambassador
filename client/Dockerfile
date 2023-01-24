# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY . ./



LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates"
LABEL alias="yourtool"


RUN go mod download


RUN go build -o /docker-gs-ping

EXPOSE 3000

CMD [ "/docker-gs-ping" ]
