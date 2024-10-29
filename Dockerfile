FROM golang:1.23.2 as builder

RUN mkdir -p $GOPATH/src
WORKDIR $GOPATH/src
ADD . .
ENV GO111MODULE=on

RUN go build -o /bin/app ./cmd/multiplayer

FROM ubuntu:24.04

COPY --from=builder /bin/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENV TZ="Europe/Moscow"
RUN apt-get update  && apt install tzdata -y

CMD ["./app"]
