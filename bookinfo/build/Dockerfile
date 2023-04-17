FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . /go/src/github.com/cloudwego/biz-demo/bookinfo
WORKDIR /go/src/github.com/cloudwego/biz-demo/bookinfo

RUN make build-linux

FROM nicolaka/netshoot
RUN mkdir -p /app && \
    chown -R nobody:nogroup /app
COPY --from=builder /go/src/github.com/cloudwego/biz-demo/bookinfo/bin/bookinfo /app
COPY --from=builder /go/src/github.com/cloudwego/biz-demo/bookinfo/conf /app/conf

USER        nobody
WORKDIR     /app
ENV         PSM=cwg.bizdemo.bookinfo HERTZ_CONF_DIR=conf HERTZ_LOG_DIR=output/log
ENTRYPOINT  ["/app/bookinfo"]
