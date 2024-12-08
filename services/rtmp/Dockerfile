FROM alpine

WORKDIR /app

COPY app .

# GIN port
EXPOSE 3200

# GRPC port
EXPOSE 3201

# RMTP port
EXPOSE 1935

# Run the binary
CMD ["./app"]
