
FROM golang:1.23-alpine


ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64


RUN apk update && apk add --no-cache git


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o Game_Mode_Usage_Web_service ./server

EXPOSE 8080


CMD ["./Game_Mode_Usage_Web_service"]
