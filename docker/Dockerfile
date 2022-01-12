FROM golang:alpine
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /workspace
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
COPY build/main /workspace/main
CMD ["main"]