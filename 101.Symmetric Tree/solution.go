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

func isSymmetric(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
        return true
    }

    queue1 := &Queue{}
    queue2 := &Queue{}

    queue1.Enqueue(root.Left)
    queue2.Enqueue(root.Right)

    for !queue1.IsEmpty() && !queue2.IsEmpty() {
        node1 := queue1.Dequeue()
        node2 := queue2.Dequeue()

        if node1 == nil && node2 == nil {
            continue
        }

        if node1 == nil || node2 == nil {
            return false
        }

        if node1.Val != node2.Val {
            return false
        }

        queue1.Enqueue(node1.Right)
        queue2.Enqueue(node2.Left)
        queue1.Enqueue(node1.Left)
        queue2.Enqueue(node2.Right)

    }

    return queue1.IsEmpty() && queue2.IsEmpty()

}
