# Multi Stage Dockerfile to compile Golang image and create Application Container Image

# Builder Image
FROM golang as builder

RUN mkdir /build

ADD . /build

WORKDIR /build/application

RUN go build -o server

# Main Container Image. The reason for choosing alpine to create lightweight image.
FROM ubuntu

RUN mkdir /app

COPY --from=builder /build/application/server /app/

EXPOSE 8383

ENTRYPOINT ["/app/server"]
