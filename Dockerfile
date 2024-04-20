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
# build
RUN go build -o main .
# run
CMD ["./main"]