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

