package encode

import (
	"container/heap"
)

type MinHeap []*Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Data byte
}

func dfs(root *Node, ans *map[byte]string, curr string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		(*ans)[root.Data] = curr 
	}
	
	dfs(root.Left, ans, curr + "0")
	dfs(root.Right, ans, curr + "1")
}

func BuildHeap(s string) *MinHeap {
	freq := make(map[rune]int)
	var curr []*Node

	for _, c := range s {
		freq[c]++
	}

	for k,v := range freq {
		curr = append(curr, &Node{Val: v, Data: byte(k)})
	}

	pq := MinHeap(curr)
	heap.Init(&pq)
	return &pq
}

func BuildHuffmanTree(pq *MinHeap) (map[byte]string, *Node) {
	for len(*pq) > 1 {
		l := heap.Pop(pq).(*Node)
		r := heap.Pop(pq).(*Node)

		new_node := &Node{Val: l.Val + r.Val, Left: l, Right: r}
		heap.Push(pq, new_node)
	}

	root := heap.Pop(pq).(*Node)
	ans := make(map[byte]string)

	dfs(root, &ans, "")
	return ans, root
}

func EncodeString(s string, codes map[byte]string) string {
	var encoded string

	for i := 0; i < len(s); i++ {
		encoded += codes[s[i]]
	}

	return  encoded
}