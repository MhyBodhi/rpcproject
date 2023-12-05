package main

import (
    "fmt"
    //"net/rpc"
    "net/rpc/jsonrpc"
    "rpcproject/message"
)

func main() {
    // 连接到RPC服务
    client, err := jsonrpc.Dial("tcp", "localhost:1234")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer client.Close()

    // 准备远程方法的参数
    args := &message.Args{7, 3}

    // 调用远程方法 Multiply
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        fmt.Println("Error calling Multiply:", err)
        return
    }
    fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

    // 调用远程方法 Divide
    var quo message.Quotient
    err = client.Call("Arith.Divide", args, &quo)
    if err != nil {
        fmt.Println("Error calling Divide:", err)
        return
    }
    fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}

