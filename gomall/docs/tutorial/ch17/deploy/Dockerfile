FROM golang:alpine

WORKDIR /opt/go

COPY . /opt/go

ENV GOPROXY=https://goproxy.cn

RUN go build -o app .

FROM alpine:latest

ENV GO_ENV=online

RUN set sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    apk --no-cache add ca-certificates

WORKDIR /opt/go

COPY --from=0 /opt/go/app ./
COPY conf /opt/go/conf

CMD ["./app"]

