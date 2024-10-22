FROM golang:1.23-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./projects/go/user

FROM alpine:latest
WORKDIR /app/
COPY ./projects/go/user/.env.local .
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]

