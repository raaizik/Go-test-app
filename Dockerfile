# syntax=docker/dockerfile:1

#FROM golang:1.18
FROM alpine
RUN apk add --update git go musl-dev
ENV GO111MODULE=on

# Set destination for COPY
WORKDIR /go
#ADD . /go/src/go-test-app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build
RUN chmod +x ./bin/Go-test-app
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/Go-test-app main.go


# Run
ENTRYPOINT ["bin/Go-test-app"]
CMD ["-h"]
