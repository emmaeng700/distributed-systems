package main

import (
	"fmt"
	"huffman_encoding/encode"
	"huffman_encoding/decode"
)

func main() {
	s := "geeksforgeeks"

	pq := encode.BuildHeap(s)
	ans, root := encode.BuildHuffmanTree(pq)
	encoded := encode.EncodeString(s, ans)
	res := decode.DecodeString(encoded, root)

	fmt.Println("Encoded", encoded)
	fmt.Println("Decoded", res)
}