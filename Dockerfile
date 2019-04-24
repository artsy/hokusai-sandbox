FROM golang:1.12.4-alpine3.9

ADD https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64 /usr/local/bin/dumb-init
RUN chmod +x /usr/local/bin/dumb-init

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o hokusai-sandbox

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]
CMD ["./hokusai-sandbox"]
