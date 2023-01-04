FROM golang:alpine as build

RUN \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk --no-cache add git

WORKDIR /go/src/app
ADD . /go/src/app
RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE="on" go build -a -installsuffix cgo -ldflags '-w -s' -o /go/bin/app/main cmd/main.go

FROM alpine:3
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk --no-cache add ca-certificates openssl tzdata curl

COPY --from=build /go/bin/app /
ADD conf/config.json /config.json

ENTRYPOINT ["./main"]

