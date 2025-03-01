FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod tidy
RUN GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
RUN goose -dir ./sql/migrations postgres "host=0.0.0.0 port=5432 dbname=shop user=username password=passwd sslmode=disable" up 
RUN go build -o /shop cmd/main.go

# CMD ["go mod download", "go build -o /shop cmd/main.go", "make migration-up"]
FROM alpine:latest
COPY --from=builder /shop /shop
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["/shop"]


# docker build -t hello-world-go .
# docker run -d -p 8080:8080 hello-world-go
# docker logs <container_id>
# docker image ls