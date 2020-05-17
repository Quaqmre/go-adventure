FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

ENTRYPOINT [ "go", "run", "main.go"]

EXPOSE 8080