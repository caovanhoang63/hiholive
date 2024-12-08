FROM golang:1.23.4-alpine3.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN echo "Downloading packages..." && \
    go mod download


RUN echo "Compiling..." && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./services/rtmp

FROM alpine

WORKDIR /app

COPY --from=build /app/app .
COPY /services/rtmp/.env /app/.env

EXPOSE 8080

# Run the binary
CMD ["./app"]
