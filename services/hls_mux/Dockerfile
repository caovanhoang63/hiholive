FROM alpine:3

RUN apk update && apk add git

RUN apk add --no-cache ffmpeg

ARG S3FS_VERSION=v1.79

RUN apk --update add fuse alpine-sdk automake autoconf libxml2-dev fuse-dev curl-dev git bash;
RUN git clone https://github.com/s3fs-fuse/s3fs-fuse.git; \
  cd s3fs-fuse; \
  git checkout tags/${S3FS_VERSION}; \
  ./autogen.sh; \
  ./configure --prefix=/usr; \
  make; \
  make install; \
  rm -rf /var/cache/apk/*;


WORKDIR /app
COPY entrypoint.sh /entrypoint.sh

RUN mkdir hls_output

COPY app .

# GIN port
EXPOSE 3000

# GRPC port
EXPOSE 3300

# RMTP port
EXPOSE 1935

# Run the binary
CMD ["/entrypoint.sh"]
