FROM alpine:3

RUN apk update && apk add git

RUN apk add --no-cache ffmpeg

WORKDIR /app

COPY app .

# GIN port
EXPOSE 3000

# GRPC port
EXPOSE 3300

# Run the binary
CMD ["./app"]
