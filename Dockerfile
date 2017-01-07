FROM golang:1.8

RUN mkdir -p /go/src/github.com/efkbook/blog-sample
WORKDIR /go/src/github.com/efkbook/blog-sample

CMD ["go", "run", "main.go"]
EXPOSE 8080
