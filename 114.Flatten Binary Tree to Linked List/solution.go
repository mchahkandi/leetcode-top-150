/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Stack []*TreeNode

func (s *Stack) Push(v *TreeNode) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *TreeNode {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func flatten(root *TreeNode)  {

    if root == nil {
        return
    }

    tree := &TreeNode{}
    node := tree

    stack := make(Stack, 0)

    stack.Push(root)

    for len(stack) > 0 {
        current := stack.Pop()

        node.Right = current
        node.Left = nil


        if current.Right != nil {
            stack.Push(current.Right)
        }

        if current.Left != nil {
            stack.Push(current.Left)
        }

        node = node.Right

    }
    root = tree.Right
}
