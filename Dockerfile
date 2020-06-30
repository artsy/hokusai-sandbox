FROM golang:1.12.4-alpine3.9

RUN apk add dumb-init git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get github.com/streadway/amqp
RUN go build src/client/client.go
RUN go build src/server/server.go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./server"]
