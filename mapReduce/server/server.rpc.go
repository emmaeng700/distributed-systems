package main

import (
	"fmt"
	"net"
	"net/rpc"
	"mapReduce/coordinator"
)

func main() {
	coord := new(coordinator.Coordinator)
	rpc.Register(coord)

	listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080...")
    rpc.Accept(listener)
}