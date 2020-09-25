FROM golang:1.13.5 AS build

WORKDIR /vote_app

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

# cgo_enabled=1的情况下进行静态编译
RUN go version && go env && mkdir -p dist \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode "external" -extldflags "-static"' -o ./dist/vote_app ./cmd/vote/main.go \
    && mkdir -p /app \
    && mv dist/* /app/

FROM alpine:latest
# 切换软件源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
    #&& apk --no-cache add ca-certificates \
    #&& apk add --no-cache bash git openssh gcc musl-dev \
    #&& apk --no-cache add bash

WORKDIR /app
COPY --from=build /app .
EXPOSE 7688
USER root
#ENV configPath=/app/conf/config.yaml
#ENV port=7688
CMD ["./vote_app"]
#docker build . -t vote_app:0.0.1
#docker run -v `pwd`/conf/config.dev.yaml:/app/conf/config.yaml --rm -it -p 7688:7688 --name vote_app vote_app:0.0.1