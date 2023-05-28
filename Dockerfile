# syntax=docker/dockerfile:1

#FROM golang:1.18
FROM alpine
RUN apk add --update git go musl-dev
ENV GO111MODULE=on
WORKDIR /go
ADD . /go/src/go-test-app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN chmod +x ./bin/Go-test-app
RUN CGO_ENABLED=0 GOOS=linux go build -o go/src/go-test-app/bin/Go-test-app main.go


ENTRYPOINT ["go/src/go-test-app/bin/Go-test-app"]
CMD ["-h"]
