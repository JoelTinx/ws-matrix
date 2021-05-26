FROM golang:1.16.4-alpine

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
WORKDIR /go/src/app
COPY . .

RUN go build ./cmd/server.go

CMD ["./server"]