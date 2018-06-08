FROM golang:1.10.2-alpine

COPY go-kit /go/src/go-kit
COPY go-kit/docker/image_build/entrypoint.sh /

RUN chmod +x /entrypoint.sh

EXPOSE 7070 7071 7072 7073

ENTRYPOINT ["/entrypoint.sh"]