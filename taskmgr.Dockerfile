FROM golang:1.11.4-alpine as builder

WORKDIR /go/src/github.com/Ankr-network/dccn-hub
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/app_dccn_taskmgr app_dccn_taskmgr/main.go

FROM scratch

COPY --from=builder /go/src/github.com/Ankr-network/dccn-hub/cmd/app_dccn_taskmgr /
COPY --from=builder /go/src/github.com/Ankr-network/dccn-hub/app_dccn_taskmgr/config.json /
CMD ["/main"]
