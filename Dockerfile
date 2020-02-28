FROM golang:1.12 AS builder
WORKDIR $GOPATH/src/github.com/hiram66/user-service
COPY . .
ENV CGO_ENABLED=0
RUN go install $GOPATH/src/github.com/hiram66/user-service/cmd

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/cmd /go/bin/cmd
COPY --from=builder /etc/passwd /etc/passwd

CMD ["/go/bin/cmd"]
