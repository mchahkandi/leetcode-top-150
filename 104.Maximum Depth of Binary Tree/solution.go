/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Queue struct {
    nodes []*TreeNode
 }

 func (q *Queue) Enqueue(node *TreeNode) {
    q.nodes = append(q.nodes,node)
 }

func (q *Queue) Dequeue() *TreeNode {
    if len(q.nodes) == 0 {
		return nil
	}
    node := q.nodes[0]
    q.nodes = q.nodes[1:]
    return node
}
func (q *Queue) IsEmpty() bool {
    return len(q.nodes) == 0
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := &Queue{}
    queue.Enqueue(root)
    depth := 0

    for !queue.IsEmpty() {
        level := len(queue.nodes)

        for level != 0 {
            node := queue.Dequeue()

            if node.Left != nil {
                queue.Enqueue(node.Left)
            }
            if node.Right != nil {
                queue.Enqueue(node.Right)
            }
            level--
        }

        depth++

    }

    return depth

}
