package main

import "fmt"
import "net/rpc"
import "net"

type Calculater struct {}

type Params struct {
	A, B int64
}

func (Calc *Calculater) Add(param *Params, result *int64) error {
	*result = param.A + param.B
	return nil
}

func (Calc *Calculater) Sub(param *Params, result *int64) error {
	*result = param.A - param.B
	return nil
}

func (Calc *Calculater) Divide(param *Params, result *int64) error {
	res := param.A / param.B

	if res == 0 {
		return fmt.Errorf("ZeroDivisionError")
	}

	*result = res
	return nil
}

func (Calc *Calculater) Multiply(param *Params, result *int64) error {
	*result = param.A * param.B
	return nil
}

func main() {
	calc := new(Calculater)
	rpc.Register(calc)

	listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080...")
    rpc.Accept(listener)
}