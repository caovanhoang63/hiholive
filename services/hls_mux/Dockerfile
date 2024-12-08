FROM alpine

RUN apk add --no-cache ffmpeg

WORKDIR /app

COPY app .

# GIN port
EXPOSE 3000

# GRPC port
EXPOSE 3001

# RMTP port
EXPOSE 1935

# Run the binary
CMD ["./app"]
