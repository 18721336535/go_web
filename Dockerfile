# base image
FROM golang:alpine
# env GOPROXY
ENV GOPROXY https://goproxy.cn,direct
# create dir
RUN mkdir /app 
# duplicate source code to app
ADD . /app/
# go to app dir
WORKDIR /app
# 编译源码
RUN go build -o main .
# 运行
CMD ["./main"]