package main

import (
    "fmt"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "rpcproject/message"
)

func main() {
    // 创建Arith类型的实例
    arith := new(message.Arith)

    // 注册Arith类型的实例，使其成为RPC服务
    rpc.Register(arith)

    // 监听在1234端口
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 1234...")

    for {
        // 等待客户端连接
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // 使用JSON进行序列化和反序列化
        go jsonrpc.ServeConn(conn)
    }
}

