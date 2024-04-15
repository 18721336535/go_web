# base image
FROM golang:alpine
# 环境变量
ENV GOPROXY https://goproxy.cn,direct
# 创建目录
RUN mkdir /app 
# 将源码复制到app目录
ADD . /app/
# 切换到app目录
WORKDIR /app
# 编译源码
RUN go build -o main .
# 运行
CMD ["./main"]