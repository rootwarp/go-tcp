#
# Build Stage
#
FROM golang:1.13.6-alpine3.11 as builder
MAINTAINER rootwarp@gmail.com

ENV GO111MODULE=on
ADD ./ /go/src/github.com/rootwarp/tcp/serv
WORKDIR /go/src/github.com/rootwarp/tcp/serv
RUN apk add --no-cache --virtual .build-deps git
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a && go install


#
# Running Stage
#
FROM alpine:3.11

RUN apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /go/src/github.com/rootwarp/tcp/serv/fullchain.pem /go/bin/fullchain.pem
COPY --from=builder /go/src/github.com/rootwarp/tcp/serv/privkey.pem /go/bin/privkey.pem
COPY --from=builder /go/bin/serv /go/bin/serv

EXPOSE 8080

ENTRYPOINT ["/go/bin/serv"]
