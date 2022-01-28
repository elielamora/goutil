FROM golang:1.18-rc-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build

ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out"]