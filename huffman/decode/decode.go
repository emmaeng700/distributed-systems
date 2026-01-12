package decode
/*
* 0 ->left
* 1 -> right
* move till you see leaf node, append, start again
*/
import "huffman_encoding/encode"

func DecodeString(encoded string, root *encode.Node) string {
	curr := root
	var res string

	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
			case '0':
				curr = curr.Left
			case '1':
				curr = curr.Right
		}

		if curr.Left == nil && curr.Right == nil {
			res += string(curr.Data)
			curr = root
		}
	}

	return  res
}
