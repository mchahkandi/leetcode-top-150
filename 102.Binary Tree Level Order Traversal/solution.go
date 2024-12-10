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

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := make(Queue, 0)

    queue.Enqueue(root)

    array := [][]int{}

    for len(queue) > 0 {
        level := len(queue)
        arr := []int{}
        for level != 0 {
            node := queue.Dequeue()

            arr = append(arr,node.Val)

            if node.Left != nil {
                queue.Enqueue(node.Left)
            }
            if node.Right != nil {
                queue.Enqueue(node.Right)
            }
            level--
        }
        array = append(array,arr)
    }
    return array
}
