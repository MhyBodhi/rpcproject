# 使用官方的Go镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将本地的代码复制到容器中的工作目录
COPY . .

# 构建Go应用程序
RUN go build -o server server.go

# 暴露应用程序运行的端口
EXPOSE 1234

# 运行应用程序
CMD ["./server"]

