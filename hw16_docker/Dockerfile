FROM golang:1.24.0-alpine3.21 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o /shop cmd/main.go

FROM alpine:latest
COPY --from=builder /shop /shop
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./shop"]


# docker build -t hello-world-go .
# docker run -d -p 8080:8080 hello-world-go
# docker logs <container_id>
# docker image ls