/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Node struct {
	Tree *TreeNode
	CurrentVal int
}

type Stack []*Node

func (s *Stack) Push(v *Node) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *Node {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make(Stack, 0)
	sum := 0

	stack.Push(&Node{Tree: root,CurrentVal: root.Val})

	for len(stack) > 0 {

		node := stack.Pop()

		if node.Tree.Left == nil && node.Tree.Right == nil {
			sum += node.CurrentVal
		}

		if node.Tree.Right != nil {
			stack.Push(&Node{Tree: node.Tree.Right, CurrentVal: (node.CurrentVal * 10) + node.Tree.Right.Val })
		}

		if node.Tree.Left != nil {
			stack.Push(&Node{Tree: node.Tree.Left, CurrentVal: (node.CurrentVal * 10) + node.Tree.Left.Val })
		}

	}

	return sum
}
