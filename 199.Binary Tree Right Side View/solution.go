/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

  type Queue []*TreeNode

 func (q *Queue) Enqueue(node *TreeNode) *Queue {
    *q = append(*q, node)
    return q
 }
 func (q *Queue) Dequeue() *TreeNode {
    if len(*q) == 0 {
		return nil
	}
    first := (*q)[0]
    *q = (*q)[1:]
    return first
 }

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    queue := make(Queue, 0)
    queue.Enqueue(root)

    array := []int{}

    for len(queue) > 0 {
        level := len(queue)
        first := level

        for level != 0 {
            node := queue.Dequeue()

            if first == level {
                array = append(array,node.Val)
            }

            if node.Right != nil {
                queue.Enqueue(node.Right)
            }
            if node.Left != nil {
                queue.Enqueue(node.Left)
            }
            level--
        }
    }
    return array
}
