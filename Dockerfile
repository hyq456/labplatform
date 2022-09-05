#基础镜像,这也是我本地的golang版本
FROM golang:1.17.1 as builder

#环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

#指定工作目录
WORKDIR /go/src/labplatform
#复制根目录内所有源码文件到工作目录下
COPY . .

#编译
RUN GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -installsuffix cgo .

#基础镜像,运行环境
FROM ubuntu:18.04 as runner

#指定工作目录
WORKDIR /go/app
#复制二进制执行文件和必要的Configs目录到工作目录
COPY --from=builder /go/src/labplatform/labplatform .
COPY --from=builder /go/src/labplatform/config ./config

#暴露的端口
EXPOSE 18888

#运行 labs.api是我的项目名称
CMD ["./labplatform"]