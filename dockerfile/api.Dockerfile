FROM golang:1.11.4-alpine as builder

WORKDIR /go/src/github.com/Ankr-network/dccn-hub
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i  -o cmd/api app-dccn-api/main.go

FROM alpine:3.7

COPY --from=builder /go/src/github.com/Ankr-network/dccn-hub/cmd/api /
CMD ["/api"]
