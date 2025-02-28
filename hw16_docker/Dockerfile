FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod tidy
RUN go build -o /shop cmd/main.go
CMD ["make migration-up"]
# CMD ["go mod download", "go build -o /shop cmd/main.go", "make migration-up"]
# FROM alpine:latest
# COPY --from=builder /shop /shop
# EXPOSE 8080
# CMD ["/shop"]


# docker build -t hello-world-go .
# docker run -d -p 8080:8080 hello-world-go
# docker logs <container_id>
# docker image ls