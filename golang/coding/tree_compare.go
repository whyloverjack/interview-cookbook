package main

import "fmt"

// TreeNode 定义一个通用的二叉树节点结构
type TreeNode struct {
	Val   interface{} // 使用接口类型以便存储各种可能的数据类型
	Left  *TreeNode
	Right *TreeNode
}

// 判断两棵二叉树是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 如果两个节点都为空，则认为它们相同
	if p == nil && q == nil {
		return true
	}

	// 如果只有一个节点为空，则认为它们不同
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// 比较节点值是否相同
	if p.Val != q.Val {
		return false
	}

	// 递归比较左右子树是否相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// 示例：创建两棵相同的二叉树
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	// 比较两棵树是否相同
	result := isSameTree(tree1, tree2)
	fmt.Println("Trees are the same:", result) // 输出：Trees are the same: true
}
