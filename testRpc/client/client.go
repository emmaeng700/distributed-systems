package main

import "net/rpc"
import "fmt"

type Params struct {
	A, B int64
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Could not connect to rpc")
		return 
	}

	defer client.Close()

	params := Params{A: 2, B: 4}
	var result int64

	err = client.Call("Calculater.Add", &params, &result)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result:", result)
}