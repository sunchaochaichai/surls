FROM golang:1.10.2-alpine

COPY surls /go/src/surls
COPY surls/docker/image_build/entrypoint.sh /

RUN chmod +x /entrypoint.sh && \
    touch /tmp/block

EXPOSE 7070 7071 7072 7073

ENTRYPOINT ["/entrypoint.sh"]