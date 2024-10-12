/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack [][2]*TreeNode

func (s *Stack) Push(v [2]*TreeNode) {
	*s = append(*s, v)
}

func (s *Stack) Pop() [2]*TreeNode {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    stack := make(Stack, 0)
    stack.Push([2]*TreeNode{p,q})

    for len(stack) > 0 {
        current := stack.Pop()

        node1 := current[0]
        node2 := current[1]


        if node1 == nil && node2 == nil {
            continue
        }

        if node1 == nil || node2 == nil {
            return false
        }

        if node1.Val != node2.Val {
            return false
        }

        stack.Push([2]*TreeNode{node1.Right, node2.Right})
        stack.Push([2]*TreeNode{node1.Left, node2.Left})

    }

    return true

}

