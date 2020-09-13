FROM golang:1.15

WORKDIR /app
COPY go.mod .
RUN go mod download \
    && go mod verify

COPY v1 v1
WORKDIR /app/v1
VOLUME ["/gen"]

RUN go generate -mod vendor ./ \
        && go test ./... \
        && chown -R $(id -u):$(id -g) /gen

ENTRYPOINT [ "bash", "-c" ]
CMD [ "cp -r gen/* /gen" ]
