#Dockerfile
FROM golang:tip-bookworm
WORKDIR /app
COPY . .
RUN go mod tidy
CMD ["go", "run", "./cmd/main.go"]
EXPOSE 8080