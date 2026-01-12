package memtable

import "testing"

func TestInsert(t *testing.T) {
    tests := []struct {
        name     string
        values   []int32
        expected []int32
    }{
        {
            name:     "single insert",
            values:   []int32{10},
            expected: []int32{10},
        },
        {
            name:     "multiple inserts ascending",
            values:   []int32{1, 2, 3, 4, 5},
            expected: []int32{1, 2, 3, 4, 5},
        },
        {
            name:     "multiple inserts descending",
            values:   []int32{5, 4, 3, 2, 1},
            expected: []int32{1, 2, 3, 4, 5},
        },
        {
            name:     "random inserts",
            values:   []int32{3, 1, 4, 1, 5, 9, 2, 6},
            expected: []int32{1, 2, 3, 4, 5, 6, 9},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tree := Constructor()
            for _, val := range tt.values {
                tree.InsertVal(val)
            }

            result := inorderTraversal(tree.Root)
            if !equal(result, tt.expected) {
                t.Errorf("Insert() = %v, want %v", result, tt.expected)
            }

            if !isBalanced(tree.Root) {
                t.Errorf("Tree is not balanced after insertions")
            }
        })
    }
}

func TestDelete(t *testing.T) {
    tests := []struct {
        name      string
        inserts   []int32
        deletes   []int32
        expected  []int32
    }{
        {
            name:     "delete single node",
            inserts:  []int32{10},
            deletes:  []int32{10},
            expected: []int32{},
        },
        {
            name:     "delete from multiple nodes",
            inserts:  []int32{1, 2, 3, 4, 5},
            deletes:  []int32{3},
            expected: []int32{1, 2, 4, 5},
        },
        {
            name:     "delete non-existent",
            inserts:  []int32{1, 2, 3},
            deletes:  []int32{4},
            expected: []int32{1, 2, 3},
        },
        {
            name:     "delete all nodes",
            inserts:  []int32{1, 2, 3, 4, 5},
            deletes:  []int32{1, 2, 3, 4, 5},
            expected: []int32{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tree := Constructor()
            
            for _, val := range tt.inserts {
                tree.InsertVal(val)
            }

            for _, val := range tt.deletes {
                tree.DeleteVal(val)
            }

            result := inorderTraversal(tree.Root)
            if !equal(result, tt.expected) {
                t.Errorf("Delete() = %v, want %v", result, tt.expected)
            }

            if tree.Root != nil && !isBalanced(tree.Root) {
                t.Errorf("Tree is not balanced after deletions")
            }
        })
    }
}

func TestSearch(t *testing.T) {
    tree := Constructor()
    values := []int32{3, 1, 4, 1, 5, 9, 2, 6}
    
    for _, val := range values {
        tree.InsertVal(val)
    }

    tests := []struct {
        name     string
        search   int32
        expected bool
    }{
        {"find existing value", 3, true},
        {"find another existing value", 9, true},
        {"find non-existing value", 7, false},
        {"find non-existing value", 0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tree.SearchVal(tt.search)
            found := result != nil
            if found != tt.expected {
                t.Errorf("SearchVal(%d) found = %v, want %v", tt.search, found, tt.expected)
            }
        })
    }
}

func TestHeight(t *testing.T) {
    tree := Constructor()
    
    if tree.Height(tree.Root) != 0 {
        t.Errorf("Height of empty tree should be 0")
    }

    tree.InsertVal(10)
    if tree.Height(tree.Root) != 1 {
        t.Errorf("Height of single node tree should be 1")
    }

    tree.InsertVal(5)
    tree.InsertVal(15)
    tree.InsertVal(3)
    tree.InsertVal(7)

    if tree.Height(tree.Root) < 1 {
        t.Errorf("Height should be at least 1 for non-empty tree")
    }
}

func TestBalance(t *testing.T) {
    tree := Constructor()
    
    values := []int32{1, 2, 3, 4, 5, 6, 7}
    for _, val := range values {
        tree.InsertVal(val)
    }

    if !isBalanced(tree.Root) {
        t.Errorf("AVL tree should remain balanced after insertions")
    }
}

func inorderTraversal(node *Node) []int32 {
    if node == nil {
        return []int32{}
    }
    
    result := []int32{}
    result = append(result, inorderTraversal(node.Left)...)
    result = append(result, node.Val)
    result = append(result, inorderTraversal(node.Right)...)
    
    return result
}

func equal(a, b []int32) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func isBalanced(node *Node) bool {
    if node == nil {
        return true
    }
    
    leftHeight := getHeight(node.Left)
    rightHeight := getHeight(node.Right)
    balance := leftHeight - rightHeight
    
    if balance < -1 || balance > 1 {
        return false
    }
    
    return isBalanced(node.Left) && isBalanced(node.Right)
}

func getHeight(node *Node) int32 {
    if node == nil {
        return 0
    }
    return node.High
}

func TestConstructor(t *testing.T) {
    tree := Constructor()
    
    if tree == nil {
        t.Errorf("Constructor should return non-nil TreeNode")
    }
    
    if tree.Root != nil {
        t.Errorf("New tree should have nil root")
    }
}