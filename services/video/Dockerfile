FROM alpine

WORKDIR /app

COPY app .

# GIN port
EXPOSE 3000

# GRPC port
EXPOSE 3300

# Run the binary
CMD ["./app"]
