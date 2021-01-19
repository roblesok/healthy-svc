FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /go/src/app

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server

FROM scratch
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .
ENTRYPOINT ["./server"]
