ARG GO_VERSION=1.13
ARG APP_PORT=3001

FROM golang:${GO_VERSION} as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# With CGO_ENABLED=0 we are disabling cgo in order to build the
# golang app statically, this means we will include all the dependencies
# once you copy this binary to the image.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o server .

# Copy binrary from builder to a new image from scratch.
FROM scratch
COPY --from=builder app/server /app/

EXPOSE ${APP_PORT}
CMD ["./app/server"]