FROM registry.cn-hangzhou.aliyuncs.com/waymondocker/golang1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o waymon_api

RUN mkdir publish \
    && cp cinema_api publish \
    && cp -r config publish

RUN ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV GIN_MODE=release
EXPOSE 8081

ENTRYPOINT ["./waymon_api"]