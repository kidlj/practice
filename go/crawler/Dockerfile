FROM golang:1.12-alpine3.9

WORKDIR /go/src/github.com/kidlj/demo/crawler
COPY . .

# use vendor
ENV GOPROXY https://proxy.golang.org
RUN go get -d -v ./...

# see https://forums.docker.com/t/standard-init-linux-go-175-exec-user-process-caused-no-such-file/20025/12
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /crawler ./breadthfirst

FROM alpine:latest

COPY --from=0 /crawler /crawler

ENTRYPOINT ["/crawler"]

CMD ["/crawler"]