/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack []struct {
	node        *TreeNode
	currentVal  int
}

func (s *Stack) Push(el struct {
	node        *TreeNode
	currentVal  int
}) {
	*s = append(*s, el)
}

func (s *Stack) Pop() struct {
	node        *TreeNode
	currentVal  int
} {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := make(Stack, 0)
	stack.Push(struct {
		node        *TreeNode
		currentVal  int
	}{node: root, currentVal: targetSum - root.Val})

	for len(stack) > 0 {
		el := stack.Pop()

		if el.currentVal == 0 && el.node.Left == nil && el.node.Right == nil {
			return true
		}


		if el.node.Right != nil {
			stack.Push(struct {
				node        *TreeNode
				currentVal  int
			}{node: el.node.Right, currentVal: el.currentVal - el.node.Right.Val})
		}


		if el.node.Left != nil {
			stack.Push(struct {
				node        *TreeNode
				currentVal  int
			}{node: el.node.Left, currentVal: el.currentVal - el.node.Left.Val})
		}
	}

	return false
}
