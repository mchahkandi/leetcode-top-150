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

func averageOfLevels(root *TreeNode) []float64 {

    queue := make(Queue, 0)

    queue.Enqueue(root)

    average := []float64{}

    for len(queue) > 0 {
        level := len(queue)
        size := level
        sum := 0
        for level != 0 {
            node := queue.Dequeue()

            sum += node.Val

            if node.Left != nil {
                queue.Enqueue(node.Left)
            }
            if node.Right != nil {
                queue.Enqueue(node.Right)
            }
            level--
        }
        avg := float64(sum) / float64(size)
        average = append(average,avg)
    }
    return average
}
