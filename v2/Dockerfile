FROM golang:1.15

RUN apt update \
        && apt install -y git make

WORKDIR /
RUN git clone https://go.googlesource.com/go goroot \
    && cd /goroot \
    && git checkout dev.go2go \ 
    && cd /goroot/src \
    && ./all.bash
    
WORKDIR /app
COPY go.mod .
RUN go mod download \
    && go mod verify

COPY v2 v2
WORKDIR /app/v2
RUN /goroot/bin/go tool go2go translate ./slide.go2 ./slide_test.go2 \
    && /goroot/bin/go test ./...