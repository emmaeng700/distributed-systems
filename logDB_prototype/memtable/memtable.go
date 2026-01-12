package memtable

import "logDB"

type Node struct {
	Val int32
	Left *Node
	Right *Node
	High int32
}

type TreeNode struct {
	Root *Node
}

func Constructor() *TreeNode {
	return &TreeNode{}
}

func (n *TreeNode) Height(node *Node) int32{
	if node == nil {
		return 0
	}
	return node.High
}

func (n *TreeNode) LeftRotate(z *Node) *Node {
	y := z.Right
	T2 := y.Left

	y.Left = z
	z.Right = T2

	z.High = 1 + logdbprototype.Max(n.Height(z.Left), n.Height(z.Right))
	y.High = 1 + logdbprototype.Max(n.Height(y.Left), n.Height(y.Right))

	return y
}

func (n *TreeNode) RightRotate (z *Node) *Node {
	y := z.Left
	T3 := y.Right

	y.Right = z
	z.Left = T3

	z.High = 1 + logdbprototype.Max(n.Height(z.Left), n.Height(z.Right))
	y.High = 1 + logdbprototype.Max(n.Height(y.Left), n.Height(y.Right))

	return y
}

func (n *TreeNode) MinValueNode(root *Node) *Node {
	curr := root

	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func (n *TreeNode) Balance(node *Node) int32 {
	if node == nil {
		return 0
	}
	return n.Height(node.Left) - n.Height(node.Right)
}

func (n *TreeNode) Insert(root *Node, val int32) *Node{
	if root == nil {
		return &Node{Val: val, High: 1}
	} else if val < root.Val {
		root.Left = n.Insert(root.Left, val)
	} else if val > root.Val{
		root.Right = n.Insert(root.Right, val)
	} else {
		return root
	}

	root.High = 1 + max(n.Height(root.Left), n.Height(root.Right))
	balance := n.Balance(root)

	if balance > 1 && val < root.Left.Val {
		return n.RightRotate(root)
	}

	if balance < -1 && val > root.Right.Val {
		return n.LeftRotate(root)
	}

	if balance > 1 && val > root.Left.Val {
		root.Left = n.LeftRotate(root.Left)
		return n.RightRotate(root)
	}

	if balance < -1 && val < root.Right.Val {
		root.Right = n.RightRotate(root.Right)
		return n.LeftRotate(root)
	}

	return root
}

func (n *TreeNode) Delete(root *Node, val int32) *Node{
	if root == nil {
		return root
	}

	if val < root.Val {
		root.Left = n.Delete(root.Left, val)
	} else if val > root.Val {
		root.Right = n.Delete(root.Right, val)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return  temp
		}

		temp := n.MinValueNode(root.Right)
		root.Val = temp.Val
		root.Right = n.Delete(root.Right, temp.Val)
	}

	if root == nil {
		return root
	}

	root.High = 1 + logdbprototype.Max(n.Height(root.Left), n.Height(root.Right))
	balance := n.Balance(root)

	if balance > 1 && n.Balance(root.Left) >= 0 {
		return n.RightRotate(root)
	}

	if balance < -1 && n.Balance(root.Right) <= 0 {
		return n.LeftRotate(root)
	}

	if balance > 1 && n.Balance(root.Left) < 0 {
		root.Left = n.LeftRotate(root.Left)
		return n.RightRotate(root)
	}

	if balance < -1 && n.Balance(root.Right) > 0 {
		root.Right = n.RightRotate(root.Right)
		return n.LeftRotate(root)
	}

	return root
}

func (n *TreeNode) Search(root *Node, val int32) *Node {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return n.Search(root.Right, val)
	}

	return n.Search(root.Left, val)
}

func (n *TreeNode) InsertVal (val int32) {
	n.Root = n.Insert(n.Root, val)
}

func (n *TreeNode) DeleteVal(val int32) {
	n.Root = n.Delete(n.Root, val)
}

func (n *TreeNode) SearchVal(val int32) *Node {
	return n.Search(n.Root, val)
}